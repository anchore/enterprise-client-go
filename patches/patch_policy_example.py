#!/usr/bin/env python

import yaml

policy_example = '''
          allowlisted_images:
          - description: description
            id: id
            image:
              type: tag
              value: value
            name: name
            registry: registry
            repository: repository
          - description: description
            id: id
            image:
              type: tag
              value: value
            name: name
            registry: registry
            repository: repository
          allowlists:
          - description: description
            id: id
            items:
            - description: description
              expires_on: 2000-01-23T04:56:07.000+00:00
              gate: gate
              id: id
              trigger_id: trigger_id
            - description: description
              expires_on: 2000-01-23T04:56:07.000+00:00
              gate: gate
              id: id
              trigger_id: trigger_id
            name: name
            version: version
          - description: description
            id: id
            items:
            - description: description
              expires_on: 2000-01-23T04:56:07.000+00:00
              gate: gate
              id: id
              trigger_id: trigger_id
            - description: description
              expires_on: 2000-01-23T04:56:07.000+00:00
              gate: gate
              id: id
              trigger_id: trigger_id
            name: name
            version: version
          denylisted_images:
          - description: description
            id: id
            image:
              type: tag
              value: value
            name: name
            registry: registry
            repository: repository
          - description: description
            id: id
            image:
              type: tag
              value: value
            name: name
            registry: registry
            repository: repository
          description: description
          id: id
          last_updated: 0.8008282
          mappings:
          - allowlist_ids:
            - allowlist_ids
            - allowlist_ids
            description: description
            id: id
            image:
              type: tag
              value: value
            name: name
            registry: registry
            repository: repository
            rule_set_ids:
            - rule_set_ids
            - rule_set_ids
          - allowlist_ids:
            - allowlist_ids
            - allowlist_ids
            description: description
            id: id
            image:
              type: tag
              value: value
            name: name
            registry: registry
            repository: repository
            rule_set_ids:
            - rule_set_ids
            - rule_set_ids
          name: name
          rule_sets:
          - artifact_type: image
            description: description
            id: id
            name: name
            rules:
            - action: GO
              description: description
              gate: gate
              id: id
              params:
              - name: name
                value: value
              - name: name
                value: value
              recommendation: recommendation
              trigger: trigger
            - action: GO
              description: description
              gate: gate
              id: id
              params:
              - name: name
                value: value
              - name: name
                value: value
              recommendation: recommendation
              trigger: trigger
            version: version
          - artifact_type: image
            description: description
            id: id
            name: name
            rules:
            - action: GO
              description: description
              gate: gate
              id: id
              params:
              - name: name
                value: value
              - name: name
                value: value
              recommendation: recommendation
              trigger: trigger
            - action: GO
              description: description
              gate: gate
              id: id
              params:
              - name: name
                value: value
              - name: name
                value: value
              recommendation: recommendation
              trigger: trigger
            version: version
          source_mappings:
          - allowlist_ids:
            - allowlist_ids
            - allowlist_ids
            description: description
            host: host
            id: id
            name: name
            repository: repository
            rule_set_ids:
            - rule_set_ids
            - rule_set_ids
          - allowlist_ids:
            - allowlist_ids
            - allowlist_ids
            description: description
            host: host
            id: id
            name: name
            repository: repository
            rule_set_ids:
            - rule_set_ids
            - rule_set_ids
          version: version
          '''

if __name__ == '__main__':
    replacement = yaml.safe_load(policy_example)
    spec = None
    with open("pkg/enterprise/api/openapi.yaml", "r") as f:
        spec = yaml.safe_load(f)
    
    spec["components"]["schemas"]["PolicyRecord"]["example"]["policy"] = replacement

    with open("pkg/enterprise/api/openapi.yaml", "w") as f:
        yaml.safe_dump(spec, f, sort_keys=False)