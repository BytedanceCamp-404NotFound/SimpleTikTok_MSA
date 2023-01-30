package minioconnect

import (
	"testing"
)


func TestMinioConnect(t *testing.T){
	// viper.AutomaticEnv()
	// viper.GetViper().AddConfigPath("/yzx/src/SimpleTikTok/BaseInterface")
	minioClient,err := MinioConnect()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Logf("minioClient ,%v",minioClient)
	_ = minioClient 

}