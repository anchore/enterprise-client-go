# The version of this client (should be in line with the highest supported engine/enterprise version
CLIENT_VERSION = 4.6.0

# where all generated code will be located
PROJECT_ROOT = pkg
CLONE_DIR = local

# OpenAPI generator version to use
OPENAPI_GENERATOR_VERSION = v5.2.1

# --- anchore enterprise references
# a git tag/branch/commit within anchore/enterprise repo
ENTERPRISE_REF = d9a5b543d292c949fd3c508ce27233aa67734a2a
ENTERPRISE_ROOT = $(PROJECT_ROOT)/enterprise
RBAC_ROOT = $(PROJECT_ROOT)/rbac
ENGINE_ROOT = $(PROJECT_ROOT)/engine
ENTERPRISE_OPENAPI_DOC = $(PROJECT_ROOT)/swagger-enterprise-$(ENTERPRISE_REF).yaml
RBAC_OPENAPI_DOC = $(PROJECT_ROOT)/swagger-rbac-$(ENTERPRISE_REF).yaml
ENGINE_OPENAPI_DOC = $(PROJECT_ROOT)/swagger-engine-$(ENTERPRISE_REF).yaml

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
				-t /local/templates \
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
generate: generate-clients patch ## generate all client code from all swagger documents

define clone
	if [ ! -d "./${CLONE_DIR}" ]; then git clone git@github.com:anchore/enterprise.git $(CLONE_DIR); fi
	if [ -d "./${CLONE_DIR}" ]; then cd ${CLONE_DIR} && git fetch; fi
	cd ${CLONE_DIR} && git checkout ${ENTERPRISE_REF}
endef

$(PROJECT_ROOT):
	mkdir -p $(PROJECT_ROOT)

$(RBAC_OPENAPI_DOC) $(ENTERPRISE_OPENAPI_DOC) $(ENGINE_OPENAPI_DOC): $(PROJECT_ROOT) ## pull the enterprise external API swagger document
	$(call clone)
	# note: the existing upstream swagger document needs to be corrected, otherwise invalid code will be generated.
	# the tr/sed cmds are a workaround for now.
	cp $(CLONE_DIR)/anchore_enterprise/swagger/enterprise_api_swagger.yaml $(ENTERPRISE_OPENAPI_DOC)
	cp $(CLONE_DIR)/anchore_enterprise/swagger/rbac_manager_swagger.yaml $(RBAC_OPENAPI_DOC)
	cp $(CLONE_DIR)/anchore_enterprise/services/api/swagger/swagger.yaml $(ENGINE_OPENAPI_DOC)

.PHONY :=
generate-clients: $(ENTERPRISE_OPENAPI_DOC) $(ENGINE_OPENAPI_DOC) $(RBAC_OPENAPI_DOC) ## generate client code for engine external API
	$(call generate_openapi_client,$(ENTERPRISE_OPENAPI_DOC),enterprise,$(ENTERPRISE_ROOT))
	$(call generate_openapi_client,$(RBAC_OPENAPI_DOC),rbac,$(RBAC_ROOT))
	$(call generate_openapi_client,$(ENGINE_OPENAPI_DOC),engine,$(ENGINE_ROOT))
	# add any tailored code via go generate
	go generate .

.PHONY :=
clean: ## remove all swagger documents and generated client code
	rm -rf $(PROJECT_ROOT)

.PHONY :=
patch: ## applies any post-generation patches
	git apply patches/compliance_binary_properties.patch

.PHONY :=
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
