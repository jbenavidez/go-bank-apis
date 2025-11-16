# Bank accounts management system APIs (wip):

The following APIs can be use to implement a banking system on which admins are able to create customer, create accounts(checking|saving), and customer are able to deposit or withdraw X amount of money

# Stack
<ul>
<li>Go</li>
<li>PostgreSQL </li>
</ul>

# Customers APIs
<ul>
<li>Get All customers:
    <ul>
    <li>Endpoint: {domain-name}/customers</li>
     <li>Method: GET</li>
    </ul>
</li>
<li>Get customer:
    <ul>
    <li>Endpoint: {domain-name}/customers/{customer-id}</li>
     <li>Method: GET</li>
    </ul>
</li>
<li>Create customer:
    <ul>
    <li>Endpoint: {domain-name}/customers</li>
     <li>Method: Post</li>
    </ul>
</li>
<li>Update customer:
    <ul>
    <li>Endpoint: {domain-name}/customers/{customer-id}</li>
     <li>Method: PUT</li>
     </ul>
</li>
</ul>

# Bank Account APIs
<ul>
<li>Get all bank accounts for a given account:
    <ul>
    <li>Endpoint: {domain-name}/customers<{id}/accounts</li>
     <li>Method: GET</li>
    </ul>
</li>
<li>Create  Bank Account:
    <ul>
    <li>Endpoint: {domain-name}/customers<{id}/accounts</li>
     <li>Method: Post</li>
    </ul>
</li>



</ul>

# Transaction API
<ul>
<li>Perform Transaction:
 
  <li>Get Bank Account:
    <ul>
    <li>Endpoint: {domain-name}/customers<{id}/accounts{account-id}</li>
     <li>Method: Get</li>
      <li>Payload: 
    ````
```
{
    "transaction_type": "withdraw",
    "amount": 100

}
```
````
      </li>
    </ul>
</li>
</li>

</ul>