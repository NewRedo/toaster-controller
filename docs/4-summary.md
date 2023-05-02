Summary
=======

We have seen how to create Custom Resource Definitions (CRDs), Custom Resources (CRs) and Controllers with the
Operator SDK. It is also possible to extend kubectl to work with our CRDs.

Hopefully this has shown you how Kubernetes is not just a container platform, but a means of managing state in a 
real-world system.

There are other ways to define and use operators such as this, in a more conventional way. There are Ansible and Helm
operators which allow you to use those tools to define your operators.

In general, operators are programs which can manage the state of kubernetes applications, for example setting up 
databases, rotating certificates, and so on. They are a way of automating the management of your applications.
