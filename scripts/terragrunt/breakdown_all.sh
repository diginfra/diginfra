#!/usr/bin/env bash

# See https://www.diginfra.khulnasoft.com/docs/iac_tools/terragrunt for usage docs

# Output terraform plans
terragrunt run-all plan -out=diginfra-plan

# Loop through plans and output diginfra JSONs
planfiles=($(find . -name "diginfra-plan" | tr '\n' ' '))
for planfile in "${planfiles[@]}"; do
  echo "Running terraform show for $planfile";
  dir=$(dirname $planfile)
  cd $dir
  terraform show -json $(basename $planfile) > diginfra-plan.json
  cd -
  diginfra breakdown --path $dir/diginfra-plan.json --format json > $dir/diginfra-out.json
  rm $planfile
done

# Run diginfra output to merge the results
jsonfiles=($(find . -name "diginfra-out.json" | tr '\n' ' '))
diginfra output --format html $(echo ${jsonfiles[@]/#/--path }) > diginfra-report.html
diginfra output --format table $(echo ${jsonfiles[@]/#/--path })
echo "Also saved HTML report in diginfra-report.html"

# Tidy up
rm $jsonfiles
