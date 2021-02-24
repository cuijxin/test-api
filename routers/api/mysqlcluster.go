package v1

import (
	"context"
	"net/http"

	"github.com/cuijxin/mysql-cluster-presslabs/pkg/app"
	"github.com/cuijxin/mysql-cluster-presslabs/pkg/e"
	"github.com/cuijxin/mysql-cluster-presslabs/pkg/kube"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetMysqlClusters(c *gin.Context) {
	appG := app.Gin{C: c}
	namespace := c.Query("namespace")
	log.Infof("namespace: %s", namespace)
	list, err := kube.Dclientset.Resource(kube.MysqlclusterRes).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_MYSQL_CLUSTER_LIST_FAIL, nil)
		return
	}
	count := len(list.Items)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_MYSQL_CLUSTER_LIST_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": list.Items,
		"total": count,
	})
}
