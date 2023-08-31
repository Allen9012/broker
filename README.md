# broker
the broker of component


## 💡 Usage

You can import `broker` using:

```go
import (
"github.com/Allen9012/broker"
mongobrocker "github.com/Allen9012/broker/mongo"
)

```

Then use one of the helpers below:

##  Client
```go


 var Client *mongobrocker.Client

func init() {
    ctx := context.Background()
    tc := &mongobrocker.Client{
    BaseComponent: broker.NewBaseComponent(),
    RealCli: mongobrocker.NewClient(ctx, &mongobrocker.Config{
    URI:         "mongodb://localhost:27017",
    MinPoolSize: 3,
    MaxPoolSize: 3000,
    }),
 }
}


```
