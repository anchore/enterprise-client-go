### Developer Notes:
Due to a bug in OpenAPITool:  [Build error if 2 properties have binary format](https://github.com/OpenAPITools/openapi-generator/issues/8165).
The generated code should be patched to prevent a breaking change to the API.

```
diff --git a/pkg/external/api_default.go b/pkg/external/api_default.go
index 646bccb..312f4e0 100644
--- a/pkg/external/api_default.go
+++ b/pkg/external/api_default.go
@@ -322,6 +322,7 @@ func (a *DefaultApiService) AddRuntimeComplianceCheck(ctx _context.Context, chec
                localVarFormFileName string
                localVarFileName     string
                localVarFileBytes    []byte
+               localVarFile         *os.File
                localVarReturnValue  RuntimeComplianceCheck
        )
 
@@ -372,7 +373,6 @@ func (a *DefaultApiService) AddRuntimeComplianceCheck(ctx _context.Context, chec
                localVarFormParams.Add("end_time", parameterToString(localVarOptionals.EndTime.Value(), ""))
        }
        localVarFormFileName = "result_file"
-       var localVarFile *os.File
        if localVarOptionals != nil && localVarOptionals.ResultFile.IsSet() {
                localVarFileOk := false
                localVarFile, localVarFileOk = localVarOptionals.ResultFile.Value().(*os.File)
@@ -385,9 +385,9 @@ func (a *DefaultApiService) AddRuntimeComplianceCheck(ctx _context.Context, chec
                localVarFileBytes = fbs
                localVarFileName = localVarFile.Name()
                localVarFile.Close()
+               localVarFile = nil
        }
        localVarFormFileName = "report_file"
-       var localVarFile *os.File
        if localVarOptionals != nil && localVarOptionals.ReportFile.IsSet() {
                localVarFileOk := false
                localVarFile, localVarFileOk = localVarOptionals.ReportFile.Value().(*os.File)

```

This patch is applied automatically:
```
diff --git a/pkg/enterprise/model_patch_user_api_key_request.go b/pkg/enterprise/model_patch_user_api_key_request.go
index cd81fbf..e5dcf5a 100644
--- a/pkg/enterprise/model_patch_user_api_key_request.go
+++ b/pkg/enterprise/model_patch_user_api_key_request.go
@@ -18,32 +18,32 @@ import (
 
 // PatchUserApiKeyRequest struct for PatchUserApiKeyRequest
 type PatchUserApiKeyRequest struct {
-       interface{} *interface{}
+       inter interface{}
 }
 
 // Unmarshal JSON data into any of the pointers in the struct
 func (dst *PatchUserApiKeyRequest) UnmarshalJSON(data []byte) error {
        var err error
        // try to unmarshal JSON data into interface{}
-       err = json.Unmarshal(data, &dst.interface{});
+       err = json.Unmarshal(data, &dst.inter);
        if err == nil {
-               jsoninterface{}, _ := json.Marshal(dst.interface{})
-               if string(jsoninterface{}) == "{}" { // empty struct
-                       dst.interface{} = nil
+               jsoninterface, _ := json.Marshal(dst.inter)
+               if string(jsoninterface) == "{}" { // empty struct
+                       dst.inter = nil
                } else {
                        return nil // data stored in dst.interface{}, return on the first match
                }
        } else {
-               dst.interface{} = nil
+               dst.inter = nil
        }
 
-       return fmt.Errorf("Data failed to match schemas in anyOf(PatchUserApiKeyRequest)")
+       return fmt.Errorf("data failed to match schemas in anyOf(PatchUserApiKeyRequest)")
 }
 
 // Marshal data from the first non-nil pointers in the struct to JSON
 func (src *PatchUserApiKeyRequest) MarshalJSON() ([]byte, error) {
-       if src.interface{} != nil {
-               return json.Marshal(&src.interface{})
+       if src.inter != nil {
+               return json.Marshal(&src.inter)
        }
 
        return nil, nil // no data in anyOf schemas
```