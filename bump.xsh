import sys

pre_version = "1.0.2"

if 1 >= len(sys.argv) or sys.argv[1] == pre_version:
    raise Exception("need to pass new version")

version = sys.argv[1]

updated_common = $(cat backend/common.go).replace(f' = "{pre_version}"', f' = "{version}"')[:-1]
updated_bump = $(cat bump.xsh).replace(f'pre_version = "{pre_version}"', f'pre_version = "{version}"')[:-1]
updated_make = $(cat Makefile).replace(f'v ?= {pre_version}', f'v ?= {version}')[:-1]

echo @(updated_common) > backend/common.go
echo @(updated_make) > Makefile
echo @(updated_bump) > bump.xsh
git add bump.xsh Makefile backend/common.go
git commit -m f"chore: bump version to {version}"
