**gomonitorcluster**

 display all the pods across all namespaces:<br/>
 ***./gomonitorcluster --kubeconfig mycluster.conf***
 
 display all the pods running in Foo namespace:<br/>
  ***./gomonitorcluster --mynamespace Foo --kubeconfig mycluster.conf***

 find for mypod across namespace:<br/>
  ***./gomonitorcluster --mypod mypodname --kubeconfig mycluster.conf***
 
 display mypod in Foo namespace:<br/>
  ***./gomonitorcluster --mypod mypodname --mynamespace Foo --kubeconfig mycluster.conf***

