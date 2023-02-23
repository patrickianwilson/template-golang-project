import jetbrains.buildServer.configs.kotlin.*
import jetbrains.buildServer.configs.kotlin.buildFeatures.commitStatusPublisher
import jetbrains.buildServer.configs.kotlin.buildFeatures.dockerSupport
import jetbrains.buildServer.configs.kotlin.buildSteps.DockerCommandStep
import jetbrains.buildServer.configs.kotlin.buildSteps.dockerCommand
import jetbrains.buildServer.configs.kotlin.buildSteps.gradle
import jetbrains.buildServer.configs.kotlin.buildSteps.script
import jetbrains.buildServer.configs.kotlin.buildSteps.GradleBuildStep
import jetbrains.buildServer.configs.kotlin.buildSteps.ScriptBuildStep

import jetbrains.buildServer.configs.kotlin.triggers.vcs
import jetbrains.buildServer.configs.kotlin.triggers.finishBuildTrigger

/*
The settings script is an entry point for defining a TeamCity
project hierarchy. The script should contain a single call to the
project() function with a Project instance or an init function as
an argument.

VcsRoots, BuildTypes, Templates, and subprojects can be
registered inside the project using the vcsRoot(), buildType(),
template(), and subProject() methods respectively.

To debug settings scripts in command-line, run the

    mvnDebug org.jetbrains.teamcity:teamcity-configs-maven-plugin:generate

command and attach your debugger to the port 8000.

To debug in IntelliJ Idea, open the 'Maven Projects' tool window (View
-> Tool Windows -> Maven Projects), find the generate task node
(Plugins -> teamcity-configs -> teamcity-configs:generate), the
'Debug' option is available in the context menu for the task.
*/

version = "2021.2"

object BuildConstants{
    const val DOCKER_REGISTRY_ID = "PROJECT_EXT_5"
}


project {

    buildType(BuildAndInternalPublish)
    buildType(CLIRelease)
}

object BuildAndInternalPublish : BuildType({
    name = "Build and Internal Publish"
    artifactRules = "build/**/* => build_artifacts"

    vcs {
        root(DslContext.settingsRoot)
    }

    steps {
        gradle {
            tasks = "clean release"
            useGradleWrapper = false
            dockerImagePlatform = GradleBuildStep.ImagePlatform.Linux
            dockerPull = true
            dockerImage = "inquest.registry.jetbrains.space/p/buildtools/buildimages/buildimage:latest"
        }
    }

    triggers {
        vcs {
        }
    }

    features {
        dockerSupport {
            loginToRegistry = on {
                dockerRegistryId = "PROJECT_EXT_5"
            }
        }
        commitStatusPublisher {
            vcsRootExtId = "${DslContext.settingsRoot.id}"
            publisher = space {
                authType = connection {
                    connectionId = "PROJECT_EXT_2"
                }
                displayName = "TeamCity"
            }
        }
    }
})

object CLIRelease : BuildType({
    name = "Release CLI"

    artifactRules = "**/* => deploy_artifacts"
    maxRunningBuilds = 1
    type = BuildTypeSettings.Type.DEPLOYMENT
    enablePersonalBuilds = false
    steps {

        script {
            name = "CLI Release"
            scriptContent = """
                echo "Environment For Build Num: ${'$'}{BUILD_NUMBER}"
                env
                
                echo "Current Directory Contents"
                ls -al
                                    
                echo "Authenticate GCloud"
                gcloud auth activate-service-account --key-file ${'$'}GOOGLE_APPLICATION_CREDENTIALS
                                    
                echo "GCloud Tool Config"
                gcloud config list
                
                gsutil cp  darwinamd64/bin/cassius gs://downloads-wilsonsinquest-com/cassius/latest/macosx/cassius
                gsutil acl ch -u AllUsers:R gs://downloads-wilsonsinquest-com/cassius/latest/macosx/cassius
                
                gsutil cp  linuxamd64/bin/cassius gs://downloads-wilsonsinquest-com/cassius/latest/linux/cassius
                gsutil acl ch -u AllUsers:R gs://downloads-wilsonsinquest-com/cassius/latest/linux/cassius
                
            """.trimIndent()
            dockerImage = "inquest.registry.jetbrains.space/p/buildtools/buildimages/buildimage:latest"
            dockerImagePlatform = ScriptBuildStep.ImagePlatform.Linux
            dockerPull = true
        }
    }

    triggers {
        finishBuildTrigger {
            buildType = "${BuildAndInternalPublish.id}"
            successfulOnly = true
        }
    }
    features {
        dockerSupport {
            loginToRegistry = on {
                dockerRegistryId = "PROJECT_EXT_5"
            }
        }
    }

    dependencies {
        dependency(BuildAndInternalPublish) {
            snapshot {
                onDependencyFailure = FailureAction.CANCEL
                onDependencyCancel = FailureAction.CANCEL
            }

            artifacts {
                artifactRules = "build_artifacts/**"
            }
        }
    }
})
