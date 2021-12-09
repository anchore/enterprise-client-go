# The version of this client (should be in line with the highest supported engine/enterprise version
CLIENT_VERSION = 4.0.0

# where all generated code will be located
PROJECT_ROOT = pkg

# OpenAPI generator version to use
OPENAPI_GENERATOR_VERSION = v4.3.1
# note: v5 introduces the new command pattern approach, splitting request and execute + generating interfaces per service.
#OPENAPI_GENERATOR_VERSION = v5.0.0-beta3

EXTAPI_CLIENT_ROOT = $(PROJECT_ROOT)/external
EXTAPI_OPENAPI_DOC = $(PROJECT_ROOT)/swagger-external-$(ENTERPRISE_REF).yaml
SPECS = specs
EXTERNAL_SPEC = $(SPECS)/external.yaml
EXTERNAL_NAME = external
EXTERNAL_PKG = $(PROJECT_ROOT)/external
FEEDS_SPEC = $(SPECS)/feeds.yaml
FEEDS_NAME = feeds
FEEDS_PKG = $(PROJECT_ROOT)/feeds
RBAC_SPEC = $(SPECS)/rbac.yaml
RBAC_NAME = rbac
RBAC_PKG = $(PROJECT_ROOT)/rbac
NOTIFICATIONS_SPEC = $(SPECS)/notifications.yaml
NOTIFICAIONTS_NAME = notifications
NOTIFICATIONS_PKG = $(PROJECT_ROOT)/notifications
REPORTS_SPEC = $(SPECS)/reports.yaml
REPORTS_NAME = reports
REPORTS_PKG = $(PROJECT_ROOT)/reports

define generate_openapi_client
	# remove previous API clients
	@ if [[ "$$OSTYPE" == "linux-gnu" ]]; then sudo rm -rf ${3}; else rm -rf ${3}; fi

	# generate the API client
	docker run \
		--rm \
		-v $${PWD}:/local \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
			generate \
				-c /local/generator.yaml \
				--additional-properties packageName=${2} \
				--additional-properties packageVersion=$(CLIENT_VERSION) \
				--http-user-agent "anchore-client/$(CLIENT_VERSION)/go" \
				-i /local/${1} \
				-o /local/${3} \
				-g go

	# keep permissions the same as the user
	@ if [[ "$$OSTYPE" == "linux-gnu" ]]; then sudo chown -R $(shell id -u):$(shell id -g) ${3}; fi

	# remove unused files (go.mod is curated manually)
	rm ${3}/{.travis.yml,git_push.sh,go.*}
endef

.PHONY :=
.DEFAULT_GOAL :=
update: clean generate ## pull all swagger definitions and generate client code

.PHONY :=
generate: external feeds rbac reports notifications ## generate all client code from all swagger documents

.PHONY :=
external: ## generate client code for enterprise external API
	$(call generate_openapi_client,$(EXTERNAL_SPEC),$(EXTERNAL_NAME),$(EXTERNAL_PKG))

.PHONY :=
feeds: ## generate client code for enterprise feeds API
	$(call generate_openapi_client,$(FEEDS_SPEC),$(FEEDS_NAME),$(FEEDS_PKG))

.PHONY :=
notifications: ## generate client code for enterprise notfications API
	$(call generate_openapi_client,$(NOTIFICATIONS_SPEC),$(NOTIFICATIONS_NAME),$(NOTIFICATIONS_PKG))

.PHONY :=
rbac: ## generate client code for enterprise rbac API
	$(call generate_openapi_client,$(RBAC_SPEC),$(RBAC_NAME),$(RBAC_PKG))

.PHONY :=
reports: ## generate client code for enterprise reports API
	$(call generate_openapi_client,$(REPORTS_SPEC),$(REPORTS_NAME),$(REPORTS_PKG))

.PHONY :=
clean: ## remove all swagger documents and generated client code
	rm -rf $(PROJECT_ROOT)

.PHONY :=
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
