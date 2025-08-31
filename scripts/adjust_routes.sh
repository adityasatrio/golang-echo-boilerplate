#!/usr/bin/env bash
set -euo pipefail

name="${1:?module name required}"

# Compute PascalCase (capitalize first letter only)
pascal=$(printf '%s' "$name" | awk '{print toupper(substr($0,1,1)) substr($0,2)}')

routes_file="internal/applications/$name/controller/${name}_routes.go"
if [[ ! -f "$routes_file" ]]; then
  routes_file="internal/applications/$name/controller/routes.go"
fi

if [[ -f "$routes_file" ]]; then
  awk -v name="$name" -v pascal="$pascal" '
  BEGIN{OFS=""}
  {
    line=$0
    if (line ~ /^func \(c \*.*Controller\) AddRoutes\(e \*echo\.Echo\)/) {
      print "func (c *", pascal, "Controller) AddRoutes(e *echo.Echo, appName string) {"
      print "\tgroup := e.Group(appName + \"/", name, "\")"
      next
    }
    if (line ~ /^[[:space:]]*e\.POST\(\"\/[^"]*\"/) { sub(/^[[:space:]]*e\.POST\(\"\/[^"]*\"/, "\tgroup.POST(\"\"", line) }
    if (line ~ /^[[:space:]]*e\.PUT\(\"\/[^"]*\/:id\"/) { sub(/^[[:space:]]*e\.PUT\(\"\/[^"]*\/:id\"/, "\tgroup.PUT(\"/:id\"", line) }
    if (line ~ /^[[:space:]]*e\.DELETE\(\"\/[^"]*\/:id\"/) { sub(/^[[:space:]]*e\.DELETE\(\"\/[^"]*\/:id\"/, "\tgroup.DELETE(\"/:id\"", line) }
    if (line ~ /^[[:space:]]*e\.GET\(\"\/[^"]*\/:id\"/) { sub(/^[[:space:]]*e\.GET\(\"\/[^"]*\/:id\"/, "\tgroup.GET(\"/:id\"", line) }
    if (line ~ /^[[:space:]]*e\.GET\(\"\/[^"]*\"/) { sub(/^[[:space:]]*e\.GET\(\"\/[^"]*\"/, "\tgroup.GET(\"\"", line) }
    print line
  }
  ' "$routes_file" > "${routes_file}.tmp" && mv "${routes_file}.tmp" "$routes_file"
fi

echo "Routes adjusted to group pattern in: $routes_file"
