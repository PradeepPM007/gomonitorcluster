**gomonitorcluster**

 find all the pods across all namespaces:<br/>
 ***./gomonitorcluster --kubeconfig mycluster.conf***<br/>
 ![1](https://user-images.githubusercontent.com/76393384/194627391-236d66cb-359b-4ce0-a741-0d5e015dbc61.png)
 
 find all the pods running in default namespace:<br/>
  ***./gomonitorcluster --mynamespace Foo --kubeconfig mycluster.conf***<br/>
  ![2](https://user-images.githubusercontent.com/76393384/194627542-aff32951-8c79-4b45-abd3-3bcbb95de65b.png)

 find for mypod across namespace:<br/>
  ***./gomonitorcluster --mypod mypodname --kubeconfig mycluster.conf***<br/>
  ![3](https://user-images.githubusercontent.com/76393384/194627754-3b9db66c-eb00-4adf-867d-e2e4b2755b22.png)
 
 find mypod in default namespace:<br/>
  ***./gomonitorcluster --mypod mypodname --mynamespace Foo --kubeconfig mycluster.conf***<br/>
  ![4](https://user-images.githubusercontent.com/76393384/194628014-ab843e09-3334-4e92-94fa-328faa4fe3c8.png)
