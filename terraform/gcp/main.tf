resource "google_bigquery_dataset" "gcp_billing" {
  dataset_id                  = format("%s%s", lower(var.identifier), "BillingData")
  friendly_name               = "${var.identifier} Billing Data"
  description                 = ""
  location                    = "US"

  labels = {}


    access {
        role          = "OWNER"
        user_by_email = var.owner_email
    }
    access {
        role          = "OWNER"
        special_group = "projectOwners"
    }
    access {
        role          = "READER"
        special_group = "projectReaders"
    }
    access {
        role          = "WRITER"
        special_group = "projectWriters"
    }
}