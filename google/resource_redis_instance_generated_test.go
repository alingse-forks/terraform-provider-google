// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccRedisInstance_redisInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckRedisInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRedisInstance_redisInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_redis_instance.cache",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"reserved_ip_range", "region"},
			},
		},
	})
}

func testAccRedisInstance_redisInstanceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_redis_instance" "cache" {
  name           = "tf-test-memory-cache%{random_suffix}"
  memory_size_gb = 1
}
`, context)
}

func TestAccRedisInstance_redisInstanceFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  BootstrapSharedTestNetwork(t, "redis-full"),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckRedisInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRedisInstance_redisInstanceFullExample(context),
			},
			{
				ResourceName:            "google_redis_instance.cache",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"reserved_ip_range", "region"},
			},
		},
	})
}

func testAccRedisInstance_redisInstanceFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_redis_instance" "cache" {
  name           = "tf-test-ha-memory-cache%{random_suffix}"
  tier           = "STANDARD_HA"
  memory_size_gb = 1

  location_id             = "us-central1-a"
  alternative_location_id = "us-central1-f"

  authorized_network = data.google_compute_network.redis-network.id

  redis_version     = "REDIS_4_0"
  display_name      = "Terraform Test Instance"
  reserved_ip_range = "192.168.0.0/29"

  labels = {
    my_key    = "my_val"
    other_key = "other_val"
  }

  maintenance_policy {
    weekly_maintenance_window {
      day = "TUESDAY"
      start_time {
        hours = 0
        minutes = 30
        seconds = 0
        nanos = 0
      }
    }
  }
}

// This example assumes this network already exists.
// The API creates a tenant network per network authorized for a
// Redis instance and that network is not deleted when the user-created
// network (authorized_network) is deleted, so this prevents issues
// with tenant network quota.
// If this network hasn't been created and you are using this example in your
// config, add an additional network resource or change
// this from "data"to "resource"
data "google_compute_network" "redis-network" {
  name = "%{network_name}"
}
`, context)
}

func TestAccRedisInstance_redisInstanceFullWithPersistenceConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  BootstrapSharedTestNetwork(t, "redis-full-persis"),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckRedisInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRedisInstance_redisInstanceFullWithPersistenceConfigExample(context),
			},
			{
				ResourceName:            "google_redis_instance.cache-persis",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"reserved_ip_range", "region"},
			},
		},
	})
}

func testAccRedisInstance_redisInstanceFullWithPersistenceConfigExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_redis_instance" "cache-persis" {
  name           = "tf-test-ha-memory-cache-persis%{random_suffix}"
  tier           = "STANDARD_HA"
  memory_size_gb = 1
  location_id             = "us-central1-a"
  alternative_location_id = "us-central1-f"

  persistence_config {
    persistence_mode = "RDB"
    rdb_snapshot_period = "TWELVE_HOURS"
  }
}
`, context)
}

func TestAccRedisInstance_redisInstancePrivateServiceExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  BootstrapSharedTestNetwork(t, "redis-private"),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckRedisInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRedisInstance_redisInstancePrivateServiceExample(context),
			},
			{
				ResourceName:            "google_redis_instance.cache",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"reserved_ip_range", "region"},
			},
		},
	})
}

func testAccRedisInstance_redisInstancePrivateServiceExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
// This example assumes this network already exists.
// The API creates a tenant network per network authorized for a
// Redis instance and that network is not deleted when the user-created
// network (authorized_network) is deleted, so this prevents issues
// with tenant network quota.
// If this network hasn't been created and you are using this example in your
// config, add an additional network resource or change
// this from "data"to "resource"
data "google_compute_network" "redis-network" {
  name = "%{network_name}"
}

resource "google_compute_global_address" "service_range" {
  name          = "address%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = data.google_compute_network.redis-network.id
}

resource "google_service_networking_connection" "private_service_connection" {
  network                 = data.google_compute_network.redis-network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.service_range.name]
}

resource "google_redis_instance" "cache" {
  name           = "tf-test-private-cache%{random_suffix}"
  tier           = "STANDARD_HA"
  memory_size_gb = 1

  location_id             = "us-central1-a"
  alternative_location_id = "us-central1-f"

  authorized_network = data.google_compute_network.redis-network.id
  connect_mode       = "PRIVATE_SERVICE_ACCESS"

  redis_version     = "REDIS_4_0"
  display_name      = "Terraform Test Instance"

  depends_on = [google_service_networking_connection.private_service_connection]

}
`, context)
}

func TestAccRedisInstance_redisInstanceMrrExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  BootstrapSharedTestNetwork(t, "redis-mrr"),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckRedisInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRedisInstance_redisInstanceMrrExample(context),
			},
			{
				ResourceName:            "google_redis_instance.cache",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"reserved_ip_range", "region"},
			},
		},
	})
}

func testAccRedisInstance_redisInstanceMrrExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_redis_instance" "cache" {
  name           = "tf-test-mrr-memory-cache%{random_suffix}"
  tier           = "STANDARD_HA"
  memory_size_gb = 5

  location_id             = "us-central1-a"
  alternative_location_id = "us-central1-f"

  authorized_network = data.google_compute_network.redis-network.id

  redis_version     = "REDIS_6_X"
  display_name      = "Terraform Test Instance"
  reserved_ip_range = "192.168.0.0/28"
  replica_count     = 5
  read_replicas_mode = "READ_REPLICAS_ENABLED"

  labels = {
    my_key    = "my_val"
    other_key = "other_val"
  }
}

// This example assumes this network already exists.
// The API creates a tenant network per network authorized for a
// Redis instance and that network is not deleted when the user-created
// network (authorized_network) is deleted, so this prevents issues
// with tenant network quota.
// If this network hasn't been created and you are using this example in your
// config, add an additional network resource or change
// this from "data"to "resource"
data "google_compute_network" "redis-network" {
  name = "%{network_name}"
}
`, context)
}

func testAccCheckRedisInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_redis_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("RedisInstance still exists at %s", url)
			}
		}

		return nil
	}
}
