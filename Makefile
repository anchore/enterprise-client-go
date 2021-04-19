# The version of this client (should be in line with the highest supported engine/enterprise version
CLIENT_VERSION = 3.0.0

# where all generated code will be located
PROJECT_ROOT = pkg
CLONE_DIR = local

# OpenAPI generator version to use
OPENAPI_GENERATOR_VERSION = v4.3.1
# note: v5 introduces the new command pattern approach, splitting request and execute + generating interfaces per service.
#OPENAPI_GENERATOR_VERSION = v5.0.0-beta3

# --- anchore enterprise references
# a git tag/branch/commit within anchore/anchore-engine repo
ENTERPRISE_REF = a3840b19a09c5e7b6040d86cb450ab3a4d43766c
EXTAPI_CLIENT_ROOT = $(PROJECT_ROOT)/external
EXTAPI_OPENAPI_DOC = $(PROJECT_ROOT)/swagger-external-$(ENTERPRISE_REF).yaml

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
update: clean generate-external-client ## pull all swagger definitions and generate client code

.PHONY :=
generate: generate-external-client ## generate all client code from all swagger documents

.PHONY :=
clone:
	if [ ! -d "./${CLONE_DIR}" ]; then git clone git@github.com:anchore/enterprise.git $(CLONE_DIR); fi
	if [ -d "./${CLONE_DIR}" ]; then cd ${CLONE_DIR} && git pull origin master; fi
	cd ${CLONE_DIR} && git checkout ${ENTERPRISE_REF}

$(EXTAPI_OPENAPI_DOC): clone ## pull the engine external API swagger document
	mkdir -p $(PROJECT_ROOT)
	# note: the existing upstream swagger document needs to be corrected, otherwise invalid code will be generated.
	# the tr/sed cmds are a workaround for now.
	cp $(CLONE_DIR)/src/anchore_enterprise/swagger/enterprise_api_swagger.yaml $(EXTAPI_OPENAPI_DOC)

.PHONY :=
generate-external-client: $(EXTAPI_OPENAPI_DOC) ## generate client code for engine external API
	$(call generate_openapi_client,$(EXTAPI_OPENAPI_DOC),external,$(EXTAPI_CLIENT_ROOT))

.PHONY :=
clean: ## remove all swagger documents and generated client code
	rm -rf $(PROJECT_ROOT)

.PHONY :=
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
