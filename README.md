# FOAC - Terraform Provider for FOAAS

FOAC (Fuck Off as Code) Terraform provider is a wrapper for the [FOAAS](https://foaas.com/) (Fuck Off as a Service) API.

# Build

Enable Go module support
```shell
$ export GO111MODULE=on
```

Build the plugin binary
```shell
$ go build -v
```

# Install

```shell
$ cp terraform-provider-foac ~/.terraform.d/plugins/
```

# Run

Using the example Terraform script below you can create an AWS EC2 Instance with some added profanities.

```hcl
data "foac_pulp" "sammy" {
	language = "English"
	from = "Samuel L. Jackson"
}

resource "aws_instance" "web" {
	ami           = "ami-f976839e"
	instance_type = "t2.micro"

	tags = {
		Name = "Server"
		Quote = "${data.foac_pulp.sammy.message}"
	}
}
```

Now you can run the script...
```shell
$ terraform apply
```

Congratulations, you can now integrate profanities into your Infrastructure as Code projects, just like you always dreamed of.

```text
An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  + aws_instance.web
      id:                           <computed>
      ami:                          "ami-f976839e"
      arn:                          <computed>
      associate_public_ip_address:  <computed>
      availability_zone:            <computed>
      cpu_core_count:               <computed>
      cpu_threads_per_core:         <computed>
      ebs_block_device.#:           <computed>
      ephemeral_block_device.#:     <computed>
      get_password_data:            "false"
      host_id:                      <computed>
      instance_state:               <computed>
      instance_type:                "t2.micro"
      ipv6_address_count:           <computed>
      ipv6_addresses.#:             <computed>
      key_name:                     <computed>
      network_interface.#:          <computed>
      network_interface_id:         <computed>
      password_data:                <computed>
      placement_group:              <computed>
      primary_network_interface_id: <computed>
      private_dns:                  <computed>
      private_ip:                   <computed>
      public_dns:                   <computed>
      public_ip:                    <computed>
      root_block_device.#:          <computed>
      security_groups.#:            <computed>
      source_dest_check:            "true"
      subnet_id:                    <computed>
      tags.%:                       "2"
      tags.Name:                    "Server"
      tags.Quote:                   "English, motherfucker, do you speak it?"
      tenancy:                      <computed>
      volume_tags.%:                <computed>
      vpc_security_group_ids.#:     <computed>


Plan: 1 to add, 0 to change, 0 to destroy.
```

# Available Data Sources

The example below is an exhaustive list of available data sources in the provider.

```hcl
data "foac_anyway" "test" {
	company = "Acme Inc."
	from = "Bob"
}

data "foac_asshole" "test" {
	from = "Bob"
}

data "foac_awesome" "test" {
	from = "Bob"
}

data "foac_back" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_bag" "test" {
	from = "Bob"
}

data "foac_ballmer" "test" {
	name = "Alice"
	company = "Acme Inc."
	from = "Bob"
}

data "foac_bday" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_because" "test" {
	from = "Bob"
}

data "foac_blackadder" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_bm" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_bucket" "test" {
	from = "Bob"
}

data "foac_bus" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_bye" "test" {
	from = "Bob"
}

data "foac_caniuse" "test" {
	tool = "Terraform"
	from = "Bob"
}

data "foac_chainsaw" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_cocksplat" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_cool" "test" {
	from = "Bob"
}

data "foac_cup" "test" {
	from = "Bob"
}

data "foac_dalton" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_deraadt" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_diabetes" "test" {
	from = "Bob"
}

data "foac_donut" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_dosomething" "test" {
	do = "Deal"
	something = "cards"
	from = "Bob"
}

data "foac_equity" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_everyone" "test" {
	from = "Bob"
}

data "foac_everything" "test" {
	from = "Bob"
}

data "foac_family" "test" {
	from = "Bob"
}

data "foac_fascinating" "test" {
	from = "Bob"
}

data "foac_field" "test" {
	name = "Alice"
	from = "Bob"
	reference = "The Bible"
}

data "foac_flying" "test" {
	from = "Bob"
}

data "foac_fts" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_fyyff" "test" {
	from = "Bob"
}

data "foac_gfy" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_give" "test" {
	from = "Bob"
}

data "foac_greed" "test" {
	noun = "Cheese"
	from = "Bob"
}

data "foac_horse" "test" {
	from = "Bob"
}

data "foac_immensity" "test" {
	from = "Bob"
}

data "foac_ing" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_keep" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_keep_calm" "test" {
	reaction = "keep Terraforming"
	from = "Bob"
}

data "foac_king" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_life" "test" {
	from = "Bob"
}

data "foac_linus" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_look" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_looking" "test" {
	from = "Bob"
}

data "foac_madison" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_maybe" "test" {
	from = "Bob"
}

data "foac_me" "test" {
	from = "Bob"
}

data "foac_mornin" "test" {
	from = "Bob"
}

data "foac_no" "test" {
	from = "Bob"
}

data "foac_nugget" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_off" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_off_with" "test" {
	behavior = "conservatism"
	from = "Bob"
}

data "foac_outside" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_particular" "test" {
	thing = "project"
	from = "Bob"
}

data "foac_pink" "test" {
	from = "Bob"
}

data "foac_problem" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_programmer" "test" {
	from = "Bob"
}

data "foac_pulp" "test" {
	language = "English"
	from = "Bob"
}

data "foac_question" "test" {
	from = "Bob"
}

data "foac_rats_arse" "test" {
	from = "Bob"
}

data "foac_retard" "test" {
	from = "Bob"
}

data "foac_ridiculous" "test" {
	from = "Bob"
}

data "foac_rtfm" "test" {
	from = "Bob"
}

data "foac_sake" "test" {
	from = "Bob"
}

data "foac_shakespeare" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_shit" "test" {
	from = "Bob"
}

data "foac_shut_up" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_single" "test" {
	from = "Bob"
}

data "foac_thanks" "test" {
	from = "Bob"
}

data "foac_that" "test" {
	from = "Bob"
}

data "foac_think" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_thinking" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_this" "test" {
	from = "Bob"
}

data "foac_thumbs" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_too" "test" {
	from = "Bob"
}

data "foac_tucker" "test" {
	from = "Bob"
}

data "foac_what" "test" {
	from = "Bob"
}

data "foac_xmas" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_yoda" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_you" "test" {
	name = "Alice"
	from = "Bob"
}

data "foac_zayn" "test" {
	from = "Bob"
}

data "foac_zero" "test" {
	from = "Bob"
}
```
