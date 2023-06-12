Session 和 Token 都是处理 Web 应用程序中用户身份验证和授权的技术。它们的具体区别如下：

1. **Session（会话）**：

   Session 是一种服务器端的技术，用于在多个请求间保持用户的状态。当用户成功登录系统后，服务器为其创建一个唯一的会话 ID（Session ID），并将其与用户信息关联。服务器将 Session ID 发送给客户端（通常以 Cookie 形式保存在用户浏览器中），并在后续的请求中重新校验以确认用户身份。

   会话技术的关键特点包括：

   - 服务器维护：Session 状态保存在服务器端，可能对服务器资源产生一定负担。
   - 使用 Cookie：通常，客户端通过 Cookie 存储 Session ID，可能受到跨站请求攻击（CSRF）等安全威胁。

2. **Token（令牌）**：

   Token 是一种客户端的技术，用于在无状态的应用程序中维护用户状态。在用户登录成功后，服务器为用户生成一个独特的加密令牌，令牌中可能包含用户信息、有效期等数据。与 Session 不同，客户端保存和管理 Token（如在本地存储或 Cookie 中）。在后续的请求中，客户端将 Token 添加到 HTTP 请求头部，服务器验证 Token 有效性后才允许访问受保护资源。

   Token 技术的关键特点包括：

   - 客户端维护：服务器不需要维护 Token 数据。
   - 无状态：Token 可以实现服务器端的无状态（Stateless）。
   - 安全性：通常，Token 在传输过程中需要使用 HTTPS 加密，而 JWT（JSON Web Token）等 Token 类型在创建时已经进行了加密。
   
Session 和 Token 的选择取决于具体的应用场景。Session 更适用于需要长时间保持状态的场景，而 Token 是 API 和移动应用所常用的用户验证技术，适用于无状态、跨域的场景。在项目实施过程中，需要根据实际需求来选择合适的技术。

# token生成方式
Token（令牌）是一种用于标识用户身份和权限的字符串。一般用于用户登录和访问授权等场景。Token 的生成方式和实现方法有多种，以下是其中两种常见的方式：

1. JWT（JSON Web Tokens）：JWT 是一种使用 JSON 对象作为 payload 的 Token，具有轻便、跨语言、可加密等特点。JWT Token 由三部分组成：Header、Payload 和 Signature，其中 Payload 和 Signature 均使用 Base64 编码方式生成，并用“.”连接。JWT 的生成流程如下：

   - 生成 Header

     Header 部分通常包含两个属性：alg 表示 Token 的加密方式（如 HS256、RS256），typ 表示 Token 类型，这里设置为 JWT。

   - 生成 Payload

     Payload 部分通常包含身份信息、权限、到期时间等信息。

   - 生成 Signature

     Signature 部分使用 Header 和 Payload，加上一个密钥，然后使用指定的哈希算法（如 HMAC-SHA256），生成签名。

   - 最终生成 JWT Token

     将 Header、Payload 和 Signature 三部分使用“.”连接，生成完整的 Token。

2. Oauth2：Oauth2 是一种开放授权协议，它允许第三方应用程序通过在用户授权后访问资源，而不必获取他们的密码。OAuth2 的生成 Token 的流程如下：

   - 用户向客户端提供用户名和密码。

   - 客户端将用户名和密码发送到授权服务器，请求访问令牌。

   - 授权服务器验证客户端提供的用户信息，如果验证通过，颁发访问令牌。

   - 客户端使用访问令牌向资源服务器请求访问受保护的资源。

   - 资源服务器需要验证访问令牌的有效性，如果验证通过，返回受保护的资源。

以上是两种常见的 Token 生成方式，Token 的生成方式和实现方法还有很多，具体应根据业务需求和安全性要求来选择。