---
title: GraphQL Query Static Rate Limiting
keywords:
  - policies
  - ratelimit
  - graphql
sidebar_position: 1
---

```mdx-code-block
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import Zoom from 'react-medium-image-zoom';
```

In this tutorial, we will use [Flow Classifier Rego Rules][rego-rules] to
statically rate limit a GraphQL query.

## Policy

We will use a policy that will rate limit unique users based on `user_id` [Flow
Label][flow-label]. This label is extracted using [Flow
Classifier][flow-classifier] and is mapped from the `userID` claim in the JWT
token sent as Authorization header in the request.

:::tip

You can quickly write classification rules on
[HTTP requests](concepts/integrations/flow-control/flow-classifier.md#live-previewing-requests)
and define scheduler priorities on
[Flow Labels](concepts/integrations/flow-control/flow-label.md#live-previewing-flow-labels)
by live previewing them first via introspection APIs.

:::

```mdx-code-block
<Tabs>
<TabItem value="YAML">
```

```yaml
{@include: ./assets/graphql-rate-limiting/graphql-rate-limiting-jwt.yaml}
```

```mdx-code-block
</TabItem>
<TabItem value="Jsonnet">
```

```jsonnet
{@include: ./assets/graphql-rate-limiting/graphql-rate-limiting-jwt.jsonnet}
```

```mdx-code-block
</TabItem>
</Tabs>
```

For example, if the mutation query is as follows

```graphql
mutation createTodo {
  createTodo(input: { text: "todo" }) {
    user {
      id
    }
    text
    done
  }
}
```

Without diving deep into how Rego works, the source section mentioned in this
tutorial does the following:

1. Parse the query
2. Check if the mutation query is `createTodo`
3. Verify the JWT token with a very secretive secret key `secret` (only for
   demonstration purposes)
4. Decode the JWT token and extract the `userID` from the claims
5. Assign the value of `userID` to the exported variable `userID` in Rego source

From there on, the classifier rule assigns the value of the exported variable
`userID` in Rego source to `user_id` flow label, effectively creating a label
`user_id:1`. This label is used by the Rate Limiter component in the policy to
limit the `createTodo` mutation query to 10 requests/second for each userID.

### Circuit Diagram

```mermaid
{@include: ./assets/graphql-rate-limiting/graphql-rate-limiting-jwt.mmd}
```

### Playground

The traffic generator for this example is configured to generate 50
requests/second for 2 minutes. When the above policy is loaded in the
playground, we see that no more than 10 requests are accepted at any given time
and rest of the requests are rejected.

<Zoom>

![GraphQL Status Rate Limiting](./assets/graphql-rate-limiting/graphql-rate-limiting-counter.png)

</Zoom>

[rego-rules]: /concepts/integrations/flow-control/flow-classifier.md#rego
[flow-label]: /concepts/integrations/flow-control/flow-label.md
[flow-classifier]: /concepts/integrations/flow-control/flow-classifier.md
