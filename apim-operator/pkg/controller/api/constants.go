package api

const (
	dockerConfig  = "docker-config"
	mgwDockerFile = "dockerfile-conf"
	swaggerVolume = "swagger-volume"

	swaggerLocation    = "/usr/wso2/swagger/"
	dockerFileLocation = "/usr/wso2/dockerfile/"
	dockerConfLocation = "/kaniko/.docker"
	dockerFile         = "dockerfile"

	mgwToolkitImgConst  = "mgwToolkitImg"
	mgwRuntimeImgConst  = "mgwRuntimeImg"
	kanikoImgConst      = "kanikoImg"
	dockerRegistryConst = "dockerRegistry"
	userNameSpaceConst  = "userNameSpace"

	wso2NameSpaceConst    = "wso2-system"
	policyConfigmap       = "policy-configmap"
	mgwConfSecretConst    = "mgw-secret"
	analyticsSecretConst  = "analytics-secret"
	dockerSecretNameConst = "docker-secret"
	controllerConfName    = "controller-config"

	usernameConst = "username"
	passwordConst = "password"

	dockerhubRegistryUrl = "https://registry-1.docker.io/"
	defaultSecurity      = "default-security"
	securityExtension    = "x-mgw-security"
)
