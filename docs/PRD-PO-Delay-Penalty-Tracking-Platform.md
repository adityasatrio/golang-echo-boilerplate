# PRD: PO Delay & Penalty Tracking Platform (MVP)

## 1. Background / Problem Statement

PO data originates in SAP. Admin manually exports to Excel, then relies on supplier-provided updates (relayed manually) to track delay and calculate penalties. This creates a bottleneck (admin as single point of data entry), no audit trail, hardcoded business rules, and fragile item/price lookups.

Goal: replace the spreadsheet with a platform where each actor inputs their own data directly, while preserving the current business logic (delay/penalty calculation method) as validated by existing process.

## 2. Goals

- Remove admin as the manual relay point for supplier and QC updates
- Preserve existing delay/penalty calculation logic (proven, accepted by business)
- Add auditability, configurability, and data integrity the spreadsheet can't provide
- Defer ERP (SAP) real-time integration to Phase 2; MVP uses manual upload
- Use existing sheet column names for MVP stakeholder demo; migrate to proper naming in Phase 2

## 3. Roles (RBAC, dynamically assignable)

| Role | Scope | Key Actions |
|---|---|---|
| Admin | Global | All actions; can act on behalf of any role (logged as "acting as") |
| Supplier PIC | Assigned line items/batches only | Update production stage progress, ETD/actual dates |
| QC (dummy: "Buyer Co.") | Assigned batches | Inspection sign-off, FI Received status |

- Roles are assignable per user, many-to-many
- Scope is at line-item/batch level, not PO level, since one PO can involve multiple suppliers
- Admin can act as any role; admin actions performed "as" another role must be logged distinctly from that role's own native actions

## 4. Data Model (conceptual)

- **PO Header** — PO Number, Posting Date, source (uploaded from SAP export)
- **Line Item** — belongs to PO Header; Item No., Expected/Ordered Qty, linked to Item Master
- **Batch/Receipt** — belongs to Line Item; Actual/Adjusted Qty, Container No., dates (original SAP due date, revised ETD, New ETD, New ETA), production stage timestamps, delay status, penalty
- **Item Master** — Item No., Description, versioned Price (effective-dated)
- **Penalty Rule Config** — admin-editable tiers (delay-day thresholds → penalty %), versioned/effective-dated

**Rationale:** the same Item No. can appear multiple times under one PO as separate batches with different qty/container — confirmed from source data (e.g., item `G040-SAY-NG` split into two batches under one PO, different container numbers and delay outcomes).

## 5. Functional Requirements — Phase 1 (MVP)

### 5.1 PO Ingestion
- Manual upload (file matching SAP export structure)
- Validation on ingest: required fields present, item codes exist in Item Master
- Flag missing price mappings on ingest instead of silently producing null/error values

### 5.2 Line Item & Batch Tracking
- Supplier PIC logs stage progress and dates per batch they're assigned to
- System computes: BALANCE (Expected Qty − Loaded Qty)
- Line-item level rollup: sum of Remaining Qty across all batches per item

### 5.3 Date Tracking (Dual-Date)
- Store both: original SAP due date AND revised ETD
- Only revised ETD drives the penalty calculation (matches current accepted process)
- Original SAP due date retained for reporting / total-slippage visibility

### 5.4 Delay Calculation
- DELAY/OK status: `New ETD > revised ETD → DELAY`
- Delay days = New ETD − revised ETD
- This baseline (revised ETD, not original due date) is the current accepted business process — not treated as a defect

### 5.5 Penalty Calculation
- Penalty rate determined by **configurable tier rules** (admin sets thresholds/rates via config, not hardcoded)
- Total Penalty = Price × Qty × Penalty Rate
- Price pulled from **versioned Item Master** (effective-dated, not a static lookup range)

### 5.6 QC / Inspection
- QC role logs FI Received status independently (not admin-mediated)

### 5.7 RBAC
- Dynamic role assignment (many-to-many user–role–scope mapping)
- Permission checks are role/permission-based, not identity-based (`has_permission(...)`, not `user.type == '...'`)

### 5.8 Audit Trail
- All date changes, stage updates, and penalty-relevant fields logged with actor, timestamp, old/new value

## 6. Phase 2 (Deferred)

- SAP API integration, replacing manual upload
- Field renaming from sheet-based names to proposed clean names (see Section 8)
- Production stage flags converted from static checkboxes to timestamped event log (enables stage-duration analytics)

## 7. Open Assumptions to Validate

- Delay is measured against revised ETD, not original SAP due date — accepted as current business process
- One PO can have multiple suppliers/PICs (confirmed)
- QC and Supplier PIC are always separate people for MVP (confirmed; RBAC design allows future flexibility)
- "NEW Due Date Item" field in the source sheet has a label/formula mismatch (label says +84 days, formula computes +70 days) — deprioritized as a standalone stored field in MVP, treated as computed only if still needed

## 8. Field Naming — MVP (Sheet-Based) vs Proposed (Phase 2)

| # | MVP Name (from sheet) | Proposed Name (Phase 2) | Entity |
|---|---|---|---|
| 1 | Document Number | po_number | PO Header |
| 2 | Posting Date PO | po_created_date | PO Header |
| 3 | Free Text | product_line_desc | PO Header |
| 4 | Due Date PO | po_header_due_date | PO Header |
| 5 | 1st batch number | so_number | PO Header / Line Item |
| 6 | Due Date Item | original_committed_date | Batch |
| 7 | revised ETD by Ratna | revised_etd (delay-calc baseline) | Batch |
| 8 | NEW Due Date Item | *(deprecated as stored field — computed only)* | Batch |
| 9 | New ETD, Upd Supplier | supplier_updated_etd | Batch |
| 10 | New ETA | calculated_eta (computed) | Batch |
| 11 | Customer/Vendor Ref. No. | supplier_name | Line Item / Batch |
| 12 | Item No. | item_code | Item Master |
| 13 | Item/Service Description | item_name | Item Master |
| 14 | Remaining Open Quantity | batch_qty | Batch |
| 15 | CBM | unit_cbm | Item Master |
| 16 | TOTAL CBM | total_cbm (computed) | Batch |
| 17 | Warehouse Code | destination_warehouse_code | Batch |
| 18 | Notes/Container Number | container_number | Batch |
| 19–26 | ASSEMBLING … LOAD (8 stage flags) | production_stage_log (event-based, Phase 2) | Batch / Stage |
| 27 | BALANCE | remaining_qty_after_load (computed) | Batch |
| 28 | FI REF | inspection_ref_number | Batch |
| 29 | Proposed Container | planned_container_number | Batch |
| 30 | DELAY/NOT | delay_status (computed) | Batch |
| 31 | DELAY IN DAYS | delay_days (computed) | Batch |
| 32 | FI REQ DATE | inspection_request_date (computed) | Batch |
| 33 | INSPECTION DATE | scheduled_inspection_date (computed) | Batch |
| 34 | FI RECEIVED (YES/NO) | inspection_confirmed (boolean) | Batch / QC |
| 35 | PENALTY | penalty_rate (from config) | Batch |
| 36 | PRICE | unit_price (from Item Master) | Item Master |
| 37 | TOTAL PENALTY | penalty_amount (computed) | Batch |
| 38 | SUPPLIER REMARKS | supplier_notes | Batch |

**Note:** MVP demo UI/labels use the left-hand column (sheet-based names) for stakeholder familiarity. Backend/database schema can already use the proposed names internally if preferred — only the demo-facing labels need to match the sheet at this stage. Confirm which layer (UI only vs UI + schema) should use sheet-based naming before implementation starts.
