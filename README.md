# Kubernetes Cloud Controller Manager for CDS
`cds-cloud-controller-manager`  用于首云k8s集群的云控制器管理器

## cloud-controller-manager 部署说明

## 参数说明

| 字段                                                         | 可选值                                       | 描述                                                         |
| ------------------------------------------------------------ | -------------------------------------------- | ------------------------------------------------------------ |
| metadata.annotations:<br />service.beta.kubernetes.io/cds-load-balancer-protocol | http \| tcp                                  | 创建的 LoadBalancer 网络协议，仅支持 http 和 tcp             |
| metadata.annotations:<br />service.beta.kubernetes.io/cds-load-balancer-size | exlarge \| large \| normal \|medium \| small | LoadBalancer 规格说明如下<br />exlarge - <br />large - <br />normal - <br />medium - <br />small - |
| metadata.annotations:<br />service.beta.kubernetes.io/cds-load-balancer-max-connection | 20000                                        | LoadBalancer 连接数设定<br />< 20000 的任何整数              |
| spec.ports.protocol                                          | TCP                                          | 目前仅支持 TCP                                               |

## 使用示例

service.yaml 

```yaml 
kind: Service
apiVersion: v1
metadata:
  name: lb-tcp
  annotations:
    service.beta.kubernetes.io/cds-load-balancer-protocol: http | tcp
    service.beta.kubernetes.io/cds-load-balancer-size: exlarge | large | normal | medium | small
    service.beta.kubernetes.io/cds-load-balancer-max-connection: 20000 
spec:
  type: LoadBalancer
  selector:
    app: ccm-nginx-example
  ports:
  - name: ccm-tcp
    protocol: TCP
    port: 80
    targetPort: 80
```













