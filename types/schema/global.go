package schema

import (
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&AmazonWorkSpaces{})
	db.AutoMigrate(&AmazonWorkSpaces_Product{})
	db.AutoMigrate(&AmazonWorkSpaces_Product_Attributes{})
	db.AutoMigrate(&AmazonWorkSpaces_Term{})
	db.AutoMigrate(&AmazonWorkSpaces_Term_Attributes{})
	db.AutoMigrate(&AmazonWorkSpaces_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonWorkSpaces_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonApiGateway{})
	db.AutoMigrate(&AmazonApiGateway_Product{})
	db.AutoMigrate(&AmazonApiGateway_Product_Attributes{})
	db.AutoMigrate(&AmazonApiGateway_Term{})
	db.AutoMigrate(&AmazonApiGateway_Term_Attributes{})
	db.AutoMigrate(&AmazonApiGateway_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonApiGateway_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonRoute53{})
	db.AutoMigrate(&AmazonRoute53_Product{})
	db.AutoMigrate(&AmazonRoute53_Product_Attributes{})
	db.AutoMigrate(&AmazonRoute53_Term{})
	db.AutoMigrate(&AmazonRoute53_Term_Attributes{})
	db.AutoMigrate(&AmazonRoute53_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonRoute53_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonMQ{})
	db.AutoMigrate(&AmazonMQ_Product{})
	db.AutoMigrate(&AmazonMQ_Product_Attributes{})
	db.AutoMigrate(&AmazonMQ_Term{})
	db.AutoMigrate(&AmazonMQ_Term_Attributes{})
	db.AutoMigrate(&AmazonMQ_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonMQ_Term_PricePerUnit{})
	db.AutoMigrate(&Translate{})
	db.AutoMigrate(&Translate_Product{})
	db.AutoMigrate(&Translate_Product_Attributes{})
	db.AutoMigrate(&Translate_Term{})
	db.AutoMigrate(&Translate_Term_Attributes{})
	db.AutoMigrate(&Translate_Term_PriceDimensions{})
	db.AutoMigrate(&Translate_Term_PricePerUnit{})
	db.AutoMigrate(&AWSCodeDeploy{})
	db.AutoMigrate(&AWSCodeDeploy_Product{})
	db.AutoMigrate(&AWSCodeDeploy_Product_Attributes{})
	db.AutoMigrate(&AWSCodeDeploy_Term{})
	db.AutoMigrate(&AWSCodeDeploy_Term_Attributes{})
	db.AutoMigrate(&AWSCodeDeploy_Term_PriceDimensions{})
	db.AutoMigrate(&AWSCodeDeploy_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonS3{})
	db.AutoMigrate(&AmazonS3_Product{})
	db.AutoMigrate(&AmazonS3_Product_Attributes{})
	db.AutoMigrate(&AmazonS3_Term{})
	db.AutoMigrate(&AmazonS3_Term_Attributes{})
	db.AutoMigrate(&AmazonS3_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonS3_Term_PricePerUnit{})
	db.AutoMigrate(&AWSDeveloperSupport{})
	db.AutoMigrate(&AWSDeveloperSupport_Product{})
	db.AutoMigrate(&AWSDeveloperSupport_Product_Attributes{})
	db.AutoMigrate(&AWSDeveloperSupport_Term{})
	db.AutoMigrate(&AWSDeveloperSupport_Term_Attributes{})
	db.AutoMigrate(&AWSDeveloperSupport_Term_PriceDimensions{})
	db.AutoMigrate(&AWSDeveloperSupport_Term_PricePerUnit{})
	db.AutoMigrate(&IngestionServiceSnowball{})
	db.AutoMigrate(&IngestionServiceSnowball_Product{})
	db.AutoMigrate(&IngestionServiceSnowball_Product_Attributes{})
	db.AutoMigrate(&IngestionServiceSnowball_Term{})
	db.AutoMigrate(&IngestionServiceSnowball_Term_Attributes{})
	db.AutoMigrate(&IngestionServiceSnowball_Term_PriceDimensions{})
	db.AutoMigrate(&IngestionServiceSnowball_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonVPC{})
	db.AutoMigrate(&AmazonVPC_Product{})
	db.AutoMigrate(&AmazonVPC_Product_Attributes{})
	db.AutoMigrate(&AmazonVPC_Term{})
	db.AutoMigrate(&AmazonVPC_Term_Attributes{})
	db.AutoMigrate(&AmazonVPC_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonVPC_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonInspector{})
	db.AutoMigrate(&AmazonInspector_Product{})
	db.AutoMigrate(&AmazonInspector_Product_Attributes{})
	db.AutoMigrate(&AmazonInspector_Term{})
	db.AutoMigrate(&AmazonInspector_Term_Attributes{})
	db.AutoMigrate(&AmazonInspector_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonInspector_Term_PricePerUnit{})
	db.AutoMigrate(&AWSStorageGateway{})
	db.AutoMigrate(&AWSStorageGateway_Product{})
	db.AutoMigrate(&AWSStorageGateway_Product_Attributes{})
	db.AutoMigrate(&AWSStorageGateway_Term{})
	db.AutoMigrate(&AWSStorageGateway_Term_Attributes{})
	db.AutoMigrate(&AWSStorageGateway_Term_PriceDimensions{})
	db.AutoMigrate(&AWSStorageGateway_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonEC2{})
	db.AutoMigrate(&AmazonEC2_Product{})
	db.AutoMigrate(&AmazonEC2_Product_Attributes{})
	db.AutoMigrate(&AmazonEC2_Term{})
	db.AutoMigrate(&AmazonEC2_Term_Attributes{})
	db.AutoMigrate(&AmazonEC2_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonEC2_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonStates{})
	db.AutoMigrate(&AmazonStates_Product{})
	db.AutoMigrate(&AmazonStates_Product_Attributes{})
	db.AutoMigrate(&AmazonStates_Term{})
	db.AutoMigrate(&AmazonStates_Term_Attributes{})
	db.AutoMigrate(&AmazonStates_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonStates_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonCloudWatch{})
	db.AutoMigrate(&AmazonCloudWatch_Product{})
	db.AutoMigrate(&AmazonCloudWatch_Product_Attributes{})
	db.AutoMigrate(&AmazonCloudWatch_Term{})
	db.AutoMigrate(&AmazonCloudWatch_Term_Attributes{})
	db.AutoMigrate(&AmazonCloudWatch_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonCloudWatch_Term_PricePerUnit{})
	db.AutoMigrate(&AWSCostExplorer{})
	db.AutoMigrate(&AWSCostExplorer_Product{})
	db.AutoMigrate(&AWSCostExplorer_Product_Attributes{})
	db.AutoMigrate(&AWSCostExplorer_Term{})
	db.AutoMigrate(&AWSCostExplorer_Term_Attributes{})
	db.AutoMigrate(&AWSCostExplorer_Term_PriceDimensions{})
	db.AutoMigrate(&AWSCostExplorer_Term_PricePerUnit{})
	db.AutoMigrate(&AWSShield{})
	db.AutoMigrate(&AWSShield_Product{})
	db.AutoMigrate(&AWSShield_Product_Attributes{})
	db.AutoMigrate(&AWSShield_Term{})
	db.AutoMigrate(&AWSShield_Term_Attributes{})
	db.AutoMigrate(&AWSShield_Term_PriceDimensions{})
	db.AutoMigrate(&AWSShield_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonChimeDialin{})
	db.AutoMigrate(&AmazonChimeDialin_Product{})
	db.AutoMigrate(&AmazonChimeDialin_Product_Attributes{})
	db.AutoMigrate(&AmazonChimeDialin_Term{})
	db.AutoMigrate(&AmazonChimeDialin_Term_Attributes{})
	db.AutoMigrate(&AmazonChimeDialin_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonChimeDialin_Term_PricePerUnit{})
	db.AutoMigrate(&AWSElementalMediaLive{})
	db.AutoMigrate(&AWSElementalMediaLive_Product{})
	db.AutoMigrate(&AWSElementalMediaLive_Product_Attributes{})
	db.AutoMigrate(&AWSElementalMediaLive_Term{})
	db.AutoMigrate(&AWSElementalMediaLive_Term_Attributes{})
	db.AutoMigrate(&AWSElementalMediaLive_Term_PriceDimensions{})
	db.AutoMigrate(&AWSElementalMediaLive_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonRedshift{})
	db.AutoMigrate(&AmazonRedshift_Product{})
	db.AutoMigrate(&AmazonRedshift_Product_Attributes{})
	db.AutoMigrate(&AmazonRedshift_Term{})
	db.AutoMigrate(&AmazonRedshift_Term_Attributes{})
	db.AutoMigrate(&AmazonRedshift_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonRedshift_Term_PricePerUnit{})
	db.AutoMigrate(&OpsWorks{})
	db.AutoMigrate(&OpsWorks_Product{})
	db.AutoMigrate(&OpsWorks_Product_Attributes{})
	db.AutoMigrate(&OpsWorks_Term{})
	db.AutoMigrate(&OpsWorks_Term_Attributes{})
	db.AutoMigrate(&OpsWorks_Term_PriceDimensions{})
	db.AutoMigrate(&OpsWorks_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonWAM{})
	db.AutoMigrate(&AmazonWAM_Product{})
	db.AutoMigrate(&AmazonWAM_Product_Attributes{})
	db.AutoMigrate(&AmazonWAM_Term{})
	db.AutoMigrate(&AmazonWAM_Term_Attributes{})
	db.AutoMigrate(&AmazonWAM_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonWAM_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonCloudDirectory{})
	db.AutoMigrate(&AmazonCloudDirectory_Product{})
	db.AutoMigrate(&AmazonCloudDirectory_Product_Attributes{})
	db.AutoMigrate(&AmazonCloudDirectory_Term{})
	db.AutoMigrate(&AmazonCloudDirectory_Term_Attributes{})
	db.AutoMigrate(&AmazonCloudDirectory_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonCloudDirectory_Term_PricePerUnit{})
	db.AutoMigrate(&AWSDatabaseMigrationSvc{})
	db.AutoMigrate(&AWSDatabaseMigrationSvc_Product{})
	db.AutoMigrate(&AWSDatabaseMigrationSvc_Product_Attributes{})
	db.AutoMigrate(&AWSDatabaseMigrationSvc_Term{})
	db.AutoMigrate(&AWSDatabaseMigrationSvc_Term_Attributes{})
	db.AutoMigrate(&AWSDatabaseMigrationSvc_Term_PriceDimensions{})
	db.AutoMigrate(&AWSDatabaseMigrationSvc_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonAppStream{})
	db.AutoMigrate(&AmazonAppStream_Product{})
	db.AutoMigrate(&AmazonAppStream_Product_Attributes{})
	db.AutoMigrate(&AmazonAppStream_Term{})
	db.AutoMigrate(&AmazonAppStream_Term_Attributes{})
	db.AutoMigrate(&AmazonAppStream_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonAppStream_Term_PricePerUnit{})
	db.AutoMigrate(&AWSDeviceFarm{})
	db.AutoMigrate(&AWSDeviceFarm_Product{})
	db.AutoMigrate(&AWSDeviceFarm_Product_Attributes{})
	db.AutoMigrate(&AWSDeviceFarm_Term{})
	db.AutoMigrate(&AWSDeviceFarm_Term_Attributes{})
	db.AutoMigrate(&AWSDeviceFarm_Term_PriceDimensions{})
	db.AutoMigrate(&AWSDeviceFarm_Term_PricePerUnit{})
	db.AutoMigrate(&AWSCodeCommit{})
	db.AutoMigrate(&AWSCodeCommit_Product{})
	db.AutoMigrate(&AWSCodeCommit_Product_Attributes{})
	db.AutoMigrate(&AWSCodeCommit_Term{})
	db.AutoMigrate(&AWSCodeCommit_Term_Attributes{})
	db.AutoMigrate(&AWSCodeCommit_Term_PriceDimensions{})
	db.AutoMigrate(&AWSCodeCommit_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonSES{})
	db.AutoMigrate(&AmazonSES_Product{})
	db.AutoMigrate(&AmazonSES_Product_Attributes{})
	db.AutoMigrate(&AmazonSES_Term{})
	db.AutoMigrate(&AmazonSES_Term_Attributes{})
	db.AutoMigrate(&AmazonSES_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonSES_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonWorkDocs{})
	db.AutoMigrate(&AmazonWorkDocs_Product{})
	db.AutoMigrate(&AmazonWorkDocs_Product_Attributes{})
	db.AutoMigrate(&AmazonWorkDocs_Term{})
	db.AutoMigrate(&AmazonWorkDocs_Term_Attributes{})
	db.AutoMigrate(&AmazonWorkDocs_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonWorkDocs_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonCognitoSync{})
	db.AutoMigrate(&AmazonCognitoSync_Product{})
	db.AutoMigrate(&AmazonCognitoSync_Product_Attributes{})
	db.AutoMigrate(&AmazonCognitoSync_Term{})
	db.AutoMigrate(&AmazonCognitoSync_Term_Attributes{})
	db.AutoMigrate(&AmazonCognitoSync_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonCognitoSync_Term_PricePerUnit{})
	db.AutoMigrate(&AWSCloudTrail{})
	db.AutoMigrate(&AWSCloudTrail_Product{})
	db.AutoMigrate(&AWSCloudTrail_Product_Attributes{})
	db.AutoMigrate(&AWSCloudTrail_Term{})
	db.AutoMigrate(&AWSCloudTrail_Term_Attributes{})
	db.AutoMigrate(&AWSCloudTrail_Term_PriceDimensions{})
	db.AutoMigrate(&AWSCloudTrail_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonCognito{})
	db.AutoMigrate(&AmazonCognito_Product{})
	db.AutoMigrate(&AmazonCognito_Product_Attributes{})
	db.AutoMigrate(&AmazonCognito_Term{})
	db.AutoMigrate(&AmazonCognito_Term_Attributes{})
	db.AutoMigrate(&AmazonCognito_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonCognito_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonSWF{})
	db.AutoMigrate(&AmazonSWF_Product{})
	db.AutoMigrate(&AmazonSWF_Product_Attributes{})
	db.AutoMigrate(&AmazonSWF_Term{})
	db.AutoMigrate(&AmazonSWF_Term_Attributes{})
	db.AutoMigrate(&AmazonSWF_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonSWF_Term_PricePerUnit{})
	db.AutoMigrate(&ElasticMapReduce{})
	db.AutoMigrate(&ElasticMapReduce_Product{})
	db.AutoMigrate(&ElasticMapReduce_Product_Attributes{})
	db.AutoMigrate(&ElasticMapReduce_Term{})
	db.AutoMigrate(&ElasticMapReduce_Term_Attributes{})
	db.AutoMigrate(&ElasticMapReduce_Term_PriceDimensions{})
	db.AutoMigrate(&ElasticMapReduce_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonConnect{})
	db.AutoMigrate(&AmazonConnect_Product{})
	db.AutoMigrate(&AmazonConnect_Product_Attributes{})
	db.AutoMigrate(&AmazonConnect_Term{})
	db.AutoMigrate(&AmazonConnect_Term_Attributes{})
	db.AutoMigrate(&AmazonConnect_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonConnect_Term_PricePerUnit{})
	db.AutoMigrate(&Transcribe{})
	db.AutoMigrate(&Transcribe_Product{})
	db.AutoMigrate(&Transcribe_Product_Attributes{})
	db.AutoMigrate(&Transcribe_Term{})
	db.AutoMigrate(&Transcribe_Term_Attributes{})
	db.AutoMigrate(&Transcribe_Term_PriceDimensions{})
	db.AutoMigrate(&Transcribe_Term_PricePerUnit{})
	db.AutoMigrate(&AWSElementalMediaStore{})
	db.AutoMigrate(&AWSElementalMediaStore_Product{})
	db.AutoMigrate(&AWSElementalMediaStore_Product_Attributes{})
	db.AutoMigrate(&AWSElementalMediaStore_Term{})
	db.AutoMigrate(&AWSElementalMediaStore_Term_Attributes{})
	db.AutoMigrate(&AWSElementalMediaStore_Term_PriceDimensions{})
	db.AutoMigrate(&AWSElementalMediaStore_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonGuardDuty{})
	db.AutoMigrate(&AmazonGuardDuty_Product{})
	db.AutoMigrate(&AmazonGuardDuty_Product_Attributes{})
	db.AutoMigrate(&AmazonGuardDuty_Term{})
	db.AutoMigrate(&AmazonGuardDuty_Term_Attributes{})
	db.AutoMigrate(&AmazonGuardDuty_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonGuardDuty_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonES{})
	db.AutoMigrate(&AmazonES_Product{})
	db.AutoMigrate(&AmazonES_Product_Attributes{})
	db.AutoMigrate(&AmazonES_Term{})
	db.AutoMigrate(&AmazonES_Term_Attributes{})
	db.AutoMigrate(&AmazonES_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonES_Term_PricePerUnit{})
	db.AutoMigrate(&AWSCodePipeline{})
	db.AutoMigrate(&AWSCodePipeline_Product{})
	db.AutoMigrate(&AWSCodePipeline_Product_Attributes{})
	db.AutoMigrate(&AWSCodePipeline_Term{})
	db.AutoMigrate(&AWSCodePipeline_Term_Attributes{})
	db.AutoMigrate(&AWSCodePipeline_Term_PriceDimensions{})
	db.AutoMigrate(&AWSCodePipeline_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonRekognition{})
	db.AutoMigrate(&AmazonRekognition_Product{})
	db.AutoMigrate(&AmazonRekognition_Product_Attributes{})
	db.AutoMigrate(&AmazonRekognition_Term{})
	db.AutoMigrate(&AmazonRekognition_Term_Attributes{})
	db.AutoMigrate(&AmazonRekognition_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonRekognition_Term_PricePerUnit{})
	db.AutoMigrate(&Awswaf{})
	db.AutoMigrate(&Awswaf_Product{})
	db.AutoMigrate(&Awswaf_Product_Attributes{})
	db.AutoMigrate(&Awswaf_Term{})
	db.AutoMigrate(&Awswaf_Term_Attributes{})
	db.AutoMigrate(&Awswaf_Term_PriceDimensions{})
	db.AutoMigrate(&Awswaf_Term_PricePerUnit{})
	db.AutoMigrate(&AWSIoT{})
	db.AutoMigrate(&AWSIoT_Product{})
	db.AutoMigrate(&AWSIoT_Product_Attributes{})
	db.AutoMigrate(&AWSIoT_Term{})
	db.AutoMigrate(&AWSIoT_Term_Attributes{})
	db.AutoMigrate(&AWSIoT_Term_PriceDimensions{})
	db.AutoMigrate(&AWSIoT_Term_PricePerUnit{})
	db.AutoMigrate(&AWSQueueService{})
	db.AutoMigrate(&AWSQueueService_Product{})
	db.AutoMigrate(&AWSQueueService_Product_Attributes{})
	db.AutoMigrate(&AWSQueueService_Term{})
	db.AutoMigrate(&AWSQueueService_Term_Attributes{})
	db.AutoMigrate(&AWSQueueService_Term_PriceDimensions{})
	db.AutoMigrate(&AWSQueueService_Term_PricePerUnit{})
	db.AutoMigrate(&AWSSecretsManager{})
	db.AutoMigrate(&AWSSecretsManager_Product{})
	db.AutoMigrate(&AWSSecretsManager_Product_Attributes{})
	db.AutoMigrate(&AWSSecretsManager_Term{})
	db.AutoMigrate(&AWSSecretsManager_Term_Attributes{})
	db.AutoMigrate(&AWSSecretsManager_Term_PriceDimensions{})
	db.AutoMigrate(&AWSSecretsManager_Term_PricePerUnit{})
	db.AutoMigrate(&AWSDirectConnect{})
	db.AutoMigrate(&AWSDirectConnect_Product{})
	db.AutoMigrate(&AWSDirectConnect_Product_Attributes{})
	db.AutoMigrate(&AWSDirectConnect_Term{})
	db.AutoMigrate(&AWSDirectConnect_Term_Attributes{})
	db.AutoMigrate(&AWSDirectConnect_Term_PriceDimensions{})
	db.AutoMigrate(&AWSDirectConnect_Term_PricePerUnit{})
	db.AutoMigrate(&AWSGreengrass{})
	db.AutoMigrate(&AWSGreengrass_Product{})
	db.AutoMigrate(&AWSGreengrass_Product_Attributes{})
	db.AutoMigrate(&AWSGreengrass_Term{})
	db.AutoMigrate(&AWSGreengrass_Term_Attributes{})
	db.AutoMigrate(&AWSGreengrass_Term_PriceDimensions{})
	db.AutoMigrate(&AWSGreengrass_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonCloudFront{})
	db.AutoMigrate(&AmazonCloudFront_Product{})
	db.AutoMigrate(&AmazonCloudFront_Product_Attributes{})
	db.AutoMigrate(&AmazonCloudFront_Term{})
	db.AutoMigrate(&AmazonCloudFront_Term_Attributes{})
	db.AutoMigrate(&AmazonCloudFront_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonCloudFront_Term_PricePerUnit{})
	db.AutoMigrate(&AWSGlue{})
	db.AutoMigrate(&AWSGlue_Product{})
	db.AutoMigrate(&AWSGlue_Product_Attributes{})
	db.AutoMigrate(&AWSGlue_Term{})
	db.AutoMigrate(&AWSGlue_Term_Attributes{})
	db.AutoMigrate(&AWSGlue_Term_PriceDimensions{})
	db.AutoMigrate(&AWSGlue_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonML{})
	db.AutoMigrate(&AmazonML_Product{})
	db.AutoMigrate(&AmazonML_Product_Attributes{})
	db.AutoMigrate(&AmazonML_Term{})
	db.AutoMigrate(&AmazonML_Term_Attributes{})
	db.AutoMigrate(&AmazonML_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonML_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonPolly{})
	db.AutoMigrate(&AmazonPolly_Product{})
	db.AutoMigrate(&AmazonPolly_Product_Attributes{})
	db.AutoMigrate(&AmazonPolly_Term{})
	db.AutoMigrate(&AmazonPolly_Term_Attributes{})
	db.AutoMigrate(&AmazonPolly_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonPolly_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonKinesisVideo{})
	db.AutoMigrate(&AmazonKinesisVideo_Product{})
	db.AutoMigrate(&AmazonKinesisVideo_Product_Attributes{})
	db.AutoMigrate(&AmazonKinesisVideo_Term{})
	db.AutoMigrate(&AmazonKinesisVideo_Term_Attributes{})
	db.AutoMigrate(&AmazonKinesisVideo_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonKinesisVideo_Term_PricePerUnit{})
	db.AutoMigrate(&AWSSupportEnterprise{})
	db.AutoMigrate(&AWSSupportEnterprise_Product{})
	db.AutoMigrate(&AWSSupportEnterprise_Product_Attributes{})
	db.AutoMigrate(&AWSSupportEnterprise_Term{})
	db.AutoMigrate(&AWSSupportEnterprise_Term_Attributes{})
	db.AutoMigrate(&AWSSupportEnterprise_Term_PriceDimensions{})
	db.AutoMigrate(&AWSSupportEnterprise_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonETS{})
	db.AutoMigrate(&AmazonETS_Product{})
	db.AutoMigrate(&AmazonETS_Product_Attributes{})
	db.AutoMigrate(&AmazonETS_Term{})
	db.AutoMigrate(&AmazonETS_Term_Attributes{})
	db.AutoMigrate(&AmazonETS_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonETS_Term_PricePerUnit{})
	db.AutoMigrate(&CloudHSM{})
	db.AutoMigrate(&CloudHSM_Product{})
	db.AutoMigrate(&CloudHSM_Product_Attributes{})
	db.AutoMigrate(&CloudHSM_Term{})
	db.AutoMigrate(&CloudHSM_Term_Attributes{})
	db.AutoMigrate(&CloudHSM_Term_PriceDimensions{})
	db.AutoMigrate(&CloudHSM_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonLex{})
	db.AutoMigrate(&AmazonLex_Product{})
	db.AutoMigrate(&AmazonLex_Product_Attributes{})
	db.AutoMigrate(&AmazonLex_Term{})
	db.AutoMigrate(&AmazonLex_Term_Attributes{})
	db.AutoMigrate(&AmazonLex_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonLex_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonSimpleDB{})
	db.AutoMigrate(&AmazonSimpleDB_Product{})
	db.AutoMigrate(&AmazonSimpleDB_Product_Attributes{})
	db.AutoMigrate(&AmazonSimpleDB_Term{})
	db.AutoMigrate(&AmazonSimpleDB_Term_Attributes{})
	db.AutoMigrate(&AmazonSimpleDB_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonSimpleDB_Term_PricePerUnit{})
	db.AutoMigrate(&AWSXRay{})
	db.AutoMigrate(&AWSXRay_Product{})
	db.AutoMigrate(&AWSXRay_Product_Attributes{})
	db.AutoMigrate(&AWSXRay_Term{})
	db.AutoMigrate(&AWSXRay_Term_Attributes{})
	db.AutoMigrate(&AWSXRay_Term_PriceDimensions{})
	db.AutoMigrate(&AWSXRay_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonSumerian{})
	db.AutoMigrate(&AmazonSumerian_Product{})
	db.AutoMigrate(&AmazonSumerian_Product_Attributes{})
	db.AutoMigrate(&AmazonSumerian_Term{})
	db.AutoMigrate(&AmazonSumerian_Term_Attributes{})
	db.AutoMigrate(&AmazonSumerian_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonSumerian_Term_PricePerUnit{})
	db.AutoMigrate(&CodeBuild{})
	db.AutoMigrate(&CodeBuild_Product{})
	db.AutoMigrate(&CodeBuild_Product_Attributes{})
	db.AutoMigrate(&CodeBuild_Term{})
	db.AutoMigrate(&CodeBuild_Term_Attributes{})
	db.AutoMigrate(&CodeBuild_Term_PriceDimensions{})
	db.AutoMigrate(&CodeBuild_Term_PricePerUnit{})
	db.AutoMigrate(&AWSSupportBusiness{})
	db.AutoMigrate(&AWSSupportBusiness_Product{})
	db.AutoMigrate(&AWSSupportBusiness_Product_Attributes{})
	db.AutoMigrate(&AWSSupportBusiness_Term{})
	db.AutoMigrate(&AWSSupportBusiness_Term_Attributes{})
	db.AutoMigrate(&AWSSupportBusiness_Term_PriceDimensions{})
	db.AutoMigrate(&AWSSupportBusiness_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonLightsail{})
	db.AutoMigrate(&AmazonLightsail_Product{})
	db.AutoMigrate(&AmazonLightsail_Product_Attributes{})
	db.AutoMigrate(&AmazonLightsail_Term{})
	db.AutoMigrate(&AmazonLightsail_Term_Attributes{})
	db.AutoMigrate(&AmazonLightsail_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonLightsail_Term_PricePerUnit{})
	db.AutoMigrate(&AWSCertificateManager{})
	db.AutoMigrate(&AWSCertificateManager_Product{})
	db.AutoMigrate(&AWSCertificateManager_Product_Attributes{})
	db.AutoMigrate(&AWSCertificateManager_Term{})
	db.AutoMigrate(&AWSCertificateManager_Term_Attributes{})
	db.AutoMigrate(&AWSCertificateManager_Term_PriceDimensions{})
	db.AutoMigrate(&AWSCertificateManager_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonNeptune{})
	db.AutoMigrate(&AmazonNeptune_Product{})
	db.AutoMigrate(&AmazonNeptune_Product_Attributes{})
	db.AutoMigrate(&AmazonNeptune_Term{})
	db.AutoMigrate(&AmazonNeptune_Term_Attributes{})
	db.AutoMigrate(&AmazonNeptune_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonNeptune_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonKinesis{})
	db.AutoMigrate(&AmazonKinesis_Product{})
	db.AutoMigrate(&AmazonKinesis_Product_Attributes{})
	db.AutoMigrate(&AmazonKinesis_Term{})
	db.AutoMigrate(&AmazonKinesis_Term_Attributes{})
	db.AutoMigrate(&AmazonKinesis_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonKinesis_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonEFS{})
	db.AutoMigrate(&AmazonEFS_Product{})
	db.AutoMigrate(&AmazonEFS_Product_Attributes{})
	db.AutoMigrate(&AmazonEFS_Term{})
	db.AutoMigrate(&AmazonEFS_Term_Attributes{})
	db.AutoMigrate(&AmazonEFS_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonEFS_Term_PricePerUnit{})
	db.AutoMigrate(&AWSElementalMediaPackage{})
	db.AutoMigrate(&AWSElementalMediaPackage_Product{})
	db.AutoMigrate(&AWSElementalMediaPackage_Product_Attributes{})
	db.AutoMigrate(&AWSElementalMediaPackage_Term{})
	db.AutoMigrate(&AWSElementalMediaPackage_Term_Attributes{})
	db.AutoMigrate(&AWSElementalMediaPackage_Term_PriceDimensions{})
	db.AutoMigrate(&AWSElementalMediaPackage_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonECS{})
	db.AutoMigrate(&AmazonECS_Product{})
	db.AutoMigrate(&AmazonECS_Product_Attributes{})
	db.AutoMigrate(&AmazonECS_Term{})
	db.AutoMigrate(&AmazonECS_Term_Attributes{})
	db.AutoMigrate(&AmazonECS_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonECS_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonSNS{})
	db.AutoMigrate(&AmazonSNS_Product{})
	db.AutoMigrate(&AmazonSNS_Product_Attributes{})
	db.AutoMigrate(&AmazonSNS_Term{})
	db.AutoMigrate(&AmazonSNS_Term_Attributes{})
	db.AutoMigrate(&AmazonSNS_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonSNS_Term_PricePerUnit{})
	db.AutoMigrate(&AWSDirectoryService{})
	db.AutoMigrate(&AWSDirectoryService_Product{})
	db.AutoMigrate(&AWSDirectoryService_Product_Attributes{})
	db.AutoMigrate(&AWSDirectoryService_Term{})
	db.AutoMigrate(&AWSDirectoryService_Term_Attributes{})
	db.AutoMigrate(&AWSDirectoryService_Term_PriceDimensions{})
	db.AutoMigrate(&AWSDirectoryService_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonEKS{})
	db.AutoMigrate(&AmazonEKS_Product{})
	db.AutoMigrate(&AmazonEKS_Product_Attributes{})
	db.AutoMigrate(&AmazonEKS_Term{})
	db.AutoMigrate(&AmazonEKS_Term_Attributes{})
	db.AutoMigrate(&AmazonEKS_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonEKS_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonWorkMail{})
	db.AutoMigrate(&AmazonWorkMail_Product{})
	db.AutoMigrate(&AmazonWorkMail_Product_Attributes{})
	db.AutoMigrate(&AmazonWorkMail_Term{})
	db.AutoMigrate(&AmazonWorkMail_Term_Attributes{})
	db.AutoMigrate(&AmazonWorkMail_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonWorkMail_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonSageMaker{})
	db.AutoMigrate(&AmazonSageMaker_Product{})
	db.AutoMigrate(&AmazonSageMaker_Product_Attributes{})
	db.AutoMigrate(&AmazonSageMaker_Term{})
	db.AutoMigrate(&AmazonSageMaker_Term_Attributes{})
	db.AutoMigrate(&AmazonSageMaker_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonSageMaker_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonECR{})
	db.AutoMigrate(&AmazonECR_Product{})
	db.AutoMigrate(&AmazonECR_Product_Attributes{})
	db.AutoMigrate(&AmazonECR_Term{})
	db.AutoMigrate(&AmazonECR_Term_Attributes{})
	db.AutoMigrate(&AmazonECR_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonECR_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonChime{})
	db.AutoMigrate(&AmazonChime_Product{})
	db.AutoMigrate(&AmazonChime_Product_Attributes{})
	db.AutoMigrate(&AmazonChime_Term{})
	db.AutoMigrate(&AmazonChime_Term_Attributes{})
	db.AutoMigrate(&AmazonChime_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonChime_Term_PricePerUnit{})
	db.AutoMigrate(&AWSEvents{})
	db.AutoMigrate(&AWSEvents_Product{})
	db.AutoMigrate(&AWSEvents_Product_Attributes{})
	db.AutoMigrate(&AWSEvents_Term{})
	db.AutoMigrate(&AWSEvents_Term_Attributes{})
	db.AutoMigrate(&AWSEvents_Term_PriceDimensions{})
	db.AutoMigrate(&AWSEvents_Term_PricePerUnit{})
	db.AutoMigrate(&SnowballExtraDays{})
	db.AutoMigrate(&SnowballExtraDays_Product{})
	db.AutoMigrate(&SnowballExtraDays_Product_Attributes{})
	db.AutoMigrate(&SnowballExtraDays_Term{})
	db.AutoMigrate(&SnowballExtraDays_Term_Attributes{})
	db.AutoMigrate(&SnowballExtraDays_Term_PriceDimensions{})
	db.AutoMigrate(&SnowballExtraDays_Term_PricePerUnit{})
	db.AutoMigrate(&Comprehend{})
	db.AutoMigrate(&Comprehend_Product{})
	db.AutoMigrate(&Comprehend_Product_Attributes{})
	db.AutoMigrate(&Comprehend_Term{})
	db.AutoMigrate(&Comprehend_Term_Attributes{})
	db.AutoMigrate(&Comprehend_Term_PriceDimensions{})
	db.AutoMigrate(&Comprehend_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonGlacier{})
	db.AutoMigrate(&AmazonGlacier_Product{})
	db.AutoMigrate(&AmazonGlacier_Product_Attributes{})
	db.AutoMigrate(&AmazonGlacier_Term{})
	db.AutoMigrate(&AmazonGlacier_Term_Attributes{})
	db.AutoMigrate(&AmazonGlacier_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonGlacier_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonDynamoDB{})
	db.AutoMigrate(&AmazonDynamoDB_Product{})
	db.AutoMigrate(&AmazonDynamoDB_Product_Attributes{})
	db.AutoMigrate(&AmazonDynamoDB_Term{})
	db.AutoMigrate(&AmazonDynamoDB_Term_Attributes{})
	db.AutoMigrate(&AmazonDynamoDB_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonDynamoDB_Term_PricePerUnit{})
	db.AutoMigrate(&Mobileanalytics{})
	db.AutoMigrate(&Mobileanalytics_Product{})
	db.AutoMigrate(&Mobileanalytics_Product_Attributes{})
	db.AutoMigrate(&Mobileanalytics_Term{})
	db.AutoMigrate(&Mobileanalytics_Term_Attributes{})
	db.AutoMigrate(&Mobileanalytics_Term_PriceDimensions{})
	db.AutoMigrate(&Mobileanalytics_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonDAX{})
	db.AutoMigrate(&AmazonDAX_Product{})
	db.AutoMigrate(&AmazonDAX_Product_Attributes{})
	db.AutoMigrate(&AmazonDAX_Term{})
	db.AutoMigrate(&AmazonDAX_Term_Attributes{})
	db.AutoMigrate(&AmazonDAX_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonDAX_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonRDS{})
	db.AutoMigrate(&AmazonRDS_Product{})
	db.AutoMigrate(&AmazonRDS_Product_Attributes{})
	db.AutoMigrate(&AmazonRDS_Term{})
	db.AutoMigrate(&AmazonRDS_Term_Attributes{})
	db.AutoMigrate(&AmazonRDS_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonRDS_Term_PricePerUnit{})
	db.AutoMigrate(&AlexaWebInfoService{})
	db.AutoMigrate(&AlexaWebInfoService_Product{})
	db.AutoMigrate(&AlexaWebInfoService_Product_Attributes{})
	db.AutoMigrate(&AlexaWebInfoService_Term{})
	db.AutoMigrate(&AlexaWebInfoService_Term_Attributes{})
	db.AutoMigrate(&AlexaWebInfoService_Term_PriceDimensions{})
	db.AutoMigrate(&AlexaWebInfoService_Term_PricePerUnit{})
	db.AutoMigrate(&Datapipeline{})
	db.AutoMigrate(&Datapipeline_Product{})
	db.AutoMigrate(&Datapipeline_Product_Attributes{})
	db.AutoMigrate(&Datapipeline_Term{})
	db.AutoMigrate(&Datapipeline_Term_Attributes{})
	db.AutoMigrate(&Datapipeline_Term_PriceDimensions{})
	db.AutoMigrate(&Datapipeline_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonKinesisAnalytics{})
	db.AutoMigrate(&AmazonKinesisAnalytics_Product{})
	db.AutoMigrate(&AmazonKinesisAnalytics_Product_Attributes{})
	db.AutoMigrate(&AmazonKinesisAnalytics_Term{})
	db.AutoMigrate(&AmazonKinesisAnalytics_Term_Attributes{})
	db.AutoMigrate(&AmazonKinesisAnalytics_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonKinesisAnalytics_Term_PricePerUnit{})
	db.AutoMigrate(&AWSConfig{})
	db.AutoMigrate(&AWSConfig_Product{})
	db.AutoMigrate(&AWSConfig_Product_Attributes{})
	db.AutoMigrate(&AWSConfig_Term{})
	db.AutoMigrate(&AWSConfig_Term_Attributes{})
	db.AutoMigrate(&AWSConfig_Term_PriceDimensions{})
	db.AutoMigrate(&AWSConfig_Term_PricePerUnit{})
	db.AutoMigrate(&AWSAppSync{})
	db.AutoMigrate(&AWSAppSync_Product{})
	db.AutoMigrate(&AWSAppSync_Product_Attributes{})
	db.AutoMigrate(&AWSAppSync_Term{})
	db.AutoMigrate(&AWSAppSync_Term_Attributes{})
	db.AutoMigrate(&AWSAppSync_Term_PriceDimensions{})
	db.AutoMigrate(&AWSAppSync_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonAthena{})
	db.AutoMigrate(&AmazonAthena_Product{})
	db.AutoMigrate(&AmazonAthena_Product_Attributes{})
	db.AutoMigrate(&AmazonAthena_Term{})
	db.AutoMigrate(&AmazonAthena_Term_Attributes{})
	db.AutoMigrate(&AmazonAthena_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonAthena_Term_PricePerUnit{})
	db.AutoMigrate(&AlexaTopSites{})
	db.AutoMigrate(&AlexaTopSites_Product{})
	db.AutoMigrate(&AlexaTopSites_Product_Attributes{})
	db.AutoMigrate(&AlexaTopSites_Term{})
	db.AutoMigrate(&AlexaTopSites_Term_Attributes{})
	db.AutoMigrate(&AlexaTopSites_Term_PriceDimensions{})
	db.AutoMigrate(&AlexaTopSites_Term_PricePerUnit{})
	db.AutoMigrate(&AWSServiceCatalog{})
	db.AutoMigrate(&AWSServiceCatalog_Product{})
	db.AutoMigrate(&AWSServiceCatalog_Product_Attributes{})
	db.AutoMigrate(&AWSServiceCatalog_Term{})
	db.AutoMigrate(&AWSServiceCatalog_Term_Attributes{})
	db.AutoMigrate(&AWSServiceCatalog_Term_PriceDimensions{})
	db.AutoMigrate(&AWSServiceCatalog_Term_PricePerUnit{})
	db.AutoMigrate(&AWSBudgets{})
	db.AutoMigrate(&AWSBudgets_Product{})
	db.AutoMigrate(&AWSBudgets_Product_Attributes{})
	db.AutoMigrate(&AWSBudgets_Term{})
	db.AutoMigrate(&AWSBudgets_Term_Attributes{})
	db.AutoMigrate(&AWSBudgets_Term_PriceDimensions{})
	db.AutoMigrate(&AWSBudgets_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonPinpoint{})
	db.AutoMigrate(&AmazonPinpoint_Product{})
	db.AutoMigrate(&AmazonPinpoint_Product_Attributes{})
	db.AutoMigrate(&AmazonPinpoint_Term{})
	db.AutoMigrate(&AmazonPinpoint_Term_Attributes{})
	db.AutoMigrate(&AmazonPinpoint_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonPinpoint_Term_PricePerUnit{})
	db.AutoMigrate(&AWSElementalMediaTailor{})
	db.AutoMigrate(&AWSElementalMediaTailor_Product{})
	db.AutoMigrate(&AWSElementalMediaTailor_Product_Attributes{})
	db.AutoMigrate(&AWSElementalMediaTailor_Term{})
	db.AutoMigrate(&AWSElementalMediaTailor_Term_Attributes{})
	db.AutoMigrate(&AWSElementalMediaTailor_Term_PriceDimensions{})
	db.AutoMigrate(&AWSElementalMediaTailor_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonKinesisFirehose{})
	db.AutoMigrate(&AmazonKinesisFirehose_Product{})
	db.AutoMigrate(&AmazonKinesisFirehose_Product_Attributes{})
	db.AutoMigrate(&AmazonKinesisFirehose_Term{})
	db.AutoMigrate(&AmazonKinesisFirehose_Term_Attributes{})
	db.AutoMigrate(&AmazonKinesisFirehose_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonKinesisFirehose_Term_PricePerUnit{})
	db.AutoMigrate(&IngestionService{})
	db.AutoMigrate(&IngestionService_Product{})
	db.AutoMigrate(&IngestionService_Product_Attributes{})
	db.AutoMigrate(&IngestionService_Term{})
	db.AutoMigrate(&IngestionService_Term_Attributes{})
	db.AutoMigrate(&IngestionService_Term_PriceDimensions{})
	db.AutoMigrate(&IngestionService_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonElastiCache{})
	db.AutoMigrate(&AmazonElastiCache_Product{})
	db.AutoMigrate(&AmazonElastiCache_Product_Attributes{})
	db.AutoMigrate(&AmazonElastiCache_Term{})
	db.AutoMigrate(&AmazonElastiCache_Term_Attributes{})
	db.AutoMigrate(&AmazonElastiCache_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonElastiCache_Term_PricePerUnit{})
	db.AutoMigrate(&Awskms{})
	db.AutoMigrate(&Awskms_Product{})
	db.AutoMigrate(&Awskms_Product_Attributes{})
	db.AutoMigrate(&Awskms_Term{})
	db.AutoMigrate(&Awskms_Term_Attributes{})
	db.AutoMigrate(&Awskms_Term_PriceDimensions{})
	db.AutoMigrate(&Awskms_Term_PricePerUnit{})
	db.AutoMigrate(&AWSElementalMediaConvert{})
	db.AutoMigrate(&AWSElementalMediaConvert_Product{})
	db.AutoMigrate(&AWSElementalMediaConvert_Product_Attributes{})
	db.AutoMigrate(&AWSElementalMediaConvert_Term{})
	db.AutoMigrate(&AWSElementalMediaConvert_Term_Attributes{})
	db.AutoMigrate(&AWSElementalMediaConvert_Term_PriceDimensions{})
	db.AutoMigrate(&AWSElementalMediaConvert_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonGameLift{})
	db.AutoMigrate(&AmazonGameLift_Product{})
	db.AutoMigrate(&AmazonGameLift_Product_Attributes{})
	db.AutoMigrate(&AmazonGameLift_Term{})
	db.AutoMigrate(&AmazonGameLift_Term_Attributes{})
	db.AutoMigrate(&AmazonGameLift_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonGameLift_Term_PricePerUnit{})
	db.AutoMigrate(&AWSLambda{})
	db.AutoMigrate(&AWSLambda_Product{})
	db.AutoMigrate(&AWSLambda_Product_Attributes{})
	db.AutoMigrate(&AWSLambda_Term{})
	db.AutoMigrate(&AWSLambda_Term_Attributes{})
	db.AutoMigrate(&AWSLambda_Term_PriceDimensions{})
	db.AutoMigrate(&AWSLambda_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonMacie{})
	db.AutoMigrate(&AmazonMacie_Product{})
	db.AutoMigrate(&AmazonMacie_Product_Attributes{})
	db.AutoMigrate(&AmazonMacie_Term{})
	db.AutoMigrate(&AmazonMacie_Term_Attributes{})
	db.AutoMigrate(&AmazonMacie_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonMacie_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonCloudSearch{})
	db.AutoMigrate(&AmazonCloudSearch_Product{})
	db.AutoMigrate(&AmazonCloudSearch_Product_Attributes{})
	db.AutoMigrate(&AmazonCloudSearch_Term{})
	db.AutoMigrate(&AmazonCloudSearch_Term_Attributes{})
	db.AutoMigrate(&AmazonCloudSearch_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonCloudSearch_Term_PricePerUnit{})
	db.AutoMigrate(&AmazonQuickSight{})
	db.AutoMigrate(&AmazonQuickSight_Product{})
	db.AutoMigrate(&AmazonQuickSight_Product_Attributes{})
	db.AutoMigrate(&AmazonQuickSight_Term{})
	db.AutoMigrate(&AmazonQuickSight_Term_Attributes{})
	db.AutoMigrate(&AmazonQuickSight_Term_PriceDimensions{})
	db.AutoMigrate(&AmazonQuickSight_Term_PricePerUnit{})
	db.AutoMigrate(&ContactCenterTelecomm{})
	db.AutoMigrate(&ContactCenterTelecomm_Product{})
	db.AutoMigrate(&ContactCenterTelecomm_Product_Attributes{})
	db.AutoMigrate(&ContactCenterTelecomm_Term{})
	db.AutoMigrate(&ContactCenterTelecomm_Term_Attributes{})
	db.AutoMigrate(&ContactCenterTelecomm_Term_PriceDimensions{})
	db.AutoMigrate(&ContactCenterTelecomm_Term_PricePerUnit{})

	return nil
}
