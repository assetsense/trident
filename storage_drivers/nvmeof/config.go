package nvmeof

const NVMeoFVersion = "1.0"

type NVMeoFConfig struct {
    Version           int `json:"version"`
    TargetAddress     string `json:"targetAddress"`
    TargetPort        string `json:"targetPort"`
    NamespaceID       string `json:"namespaceID"`
    TransportType     string `json:"transportType"`
    DiscoveryProtocol string `json:"discoveryProtocol"`
	LocalConfigFilePath string `json:"localConfigFilePath"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *NVMeoFConfig) CheckForCRDControllerForbiddenAttributes() []string {
    var forbiddenAttributes []string
    if c.LocalConfigFilePath != "" {
        forbiddenAttributes = append(forbiddenAttributes, "localConfigFilePath")
    }
    return forbiddenAttributes
}

func (c *NVMeoFConfig) ExtractSecrets() map[string]string {
    secrets := make(map[string]string)

    if c.Password != "" {
        secrets["password"] = c.Password
        c.Password = "<REDACTED>" 
    }
    return secrets
}

func (c *NVMeoFConfig) UpdateSecrets(secrets map[string]string) {
    if password, ok := secrets["password"]; ok {
        c.Password = password
    }
}

    func (c *NVMeoFConfig) String() string{
		return ""
	}
	func (c *NVMeoFConfig) GoString() string{
		return ""
	}
	func (c *NVMeoFConfig) GetCredentials() (string, string, error){
		return "", "", nil
	}
	func (c *NVMeoFConfig) HasCredentials() bool{
		return false
	}
	func (c *NVMeoFConfig) SetBackendName(backendName string){
		
	}
	func (c *NVMeoFConfig) InjectSecrets(secretMap map[string]string) error{
		return nil
	}
	func (c *NVMeoFConfig) ResetSecrets(){
		
	}
	func (c *NVMeoFConfig) HideSensitiveWithSecretName(secretName string){
		
	}
	func (c *NVMeoFConfig) GetAndHideSensitive(contextKey string) map[string]string {
		sensitiveInfo := make(map[string]string)

    switch contextKey {
    case "credentials":
        sensitiveInfo["username"] = c.Username
        sensitiveInfo["password"] = c.Password

		c.Username = "*****"
        c.Password = "*****"
    }

    return sensitiveInfo
	}
	func (c *NVMeoFConfig) SpecOnlyValidation() error{
		return nil
	}