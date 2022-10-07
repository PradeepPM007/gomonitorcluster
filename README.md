gomonitorcluster

 display all the pods across all namespaces:
 **./gomonitorcluster --kubeconfig mycluster.conf**
 
 display all the pods running in Foo namespace:
  **./gomonitorcluster --namespcae Foo --kubeconfig mycluster.conf**

 find for mypod across namespace:
  **./gomonitorcluster --name mypod --kubeconfig mycluster.conf**
 
 display mypod in Foo namespace:
  **./gomonitorcluster --name mypod --namespcae Foo --kubeconfig mycluster.conf**

