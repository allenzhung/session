[
  {
    "id" : "1",
    "session_data" : "MTU2NTQwNTE2N3xEdi1CQkFFQ180SUFBUkFCRUFBQUhQLUNBQUVHYzNSeWFXNW5EQWNBQldOdmRXNTBBMmx1ZEFRQ0FBQT18-iyVwQM7SIw0jIFUZd3WhhghW1BUb_sa2L92qvun-P0=",
    "created_on" : "2019-08-10 10:46:07.052012+08:00",
    "modified_on" : "2019-08-10 10:46:07.052012+08:00",
    "expires_on" : "2019-08-10 11:46:07.052014+08:00"
  }
]

--------------------------------------------------------------------
Mac vscode 导航 是option+command+ ←/→
iMac 自定义 搜索go back go forward ctrl+shift+←/→
ctrl+tab 切换两个文件 （菜单-转到-切换编辑器-组中下一个使用过的编辑器）

--------------------------------------------------------------------
一、程序启动初始化
store := sqlitecookie.NewStore([]byte("secret"))
返回store 为sqlitecookie.Store接口类型，
type Store interface {
	sessions.Store      
  }
}  
type Store interface {
  gsessions.Store    需要实现Get(), New(), Save()方法
  Options(Options)   需要实现Option()发发 
}
实例是type store struct {
  *sqlitestore.SqliteStore 
     这也是一个结构   实现了很多方法
     Close()
     Get()
     New()
     Save()
     insert()
     Delete()
     save()
     load()
}
结构store另外定义了Option()方法

#110
Codecs:     securecookie.CodecsFromPairs(keyPairs...),  返回[]Codec
      codecs[i/2] = New(keyPairs[i], blockKey)

type Codec interface { //接口，要求实现Encode() Decode() 方法
	Encode(name string, value interface{}) (string, error)
	Decode(name, value string, dst interface{}) error
}

func New(hashKey, blockKey []byte) *SecureCookie {
  SecureCookie结构实现了Encode() Decode() 方法
  func (s *SecureCookie) Encode(name string, value interface{}) (string, error) {
    5步， 系列化，加密，hmac/Mac，转base64，长度检查
  func (s *SecureCookie) Decode(name, value string, dst interface{}) error {    

------------------------------------------------------------------
二、每次浏览器请求访问时，中间件，设置DefaultKey 构建一个session
1、
r.Use(sessions.Sessions("mysession", store))中有设置DefaultKey
这里的store上面已说，函数要求类型是
type Store interface {
	gsessions.Store
	Options(Options)
}，实际类型是是接口
type Store interface {
	sessions.Store      //这个接口嵌套的就是上面的
}
 
2、设置DefaultKey
#63 c.Set(DefaultKey, s)
    s为 session 结构
    type session struct {
      name    string
      request *http.Request
      store   Store    //实例是前面NewStore给出的
      session *gsessions.Session  //初始nil
      written bool
      writer  http.ResponseWriter
    }
    s中的store是接口类型 type Store interface {
      sessions.Store
    }，
    但是实例是前面NewStore给出的
 
--------------------------------------------------------------------
三、
session := sessions.Default(c)
先看上下文中是否有DefaultKey
DefaultKey  = "github.com/gin-contrib/sessions"
有的话，取出session结构，断言为Session接口

--------------------------------------------------------------------
四、获取gsession中Session 结构中的Values map 中的count的值
v := session.Get("count")
     return s.Session().Values[key]
          s.session, err = s.store.Get(s.request, s.name)
          store的实例中有sqlitestore.SqliteStore结构， Get方法是它提供的
          
func (m *SqliteStore) Get(r *http.Request, name string) (*sessions.Session, error) {
            return sessions.GetRegistry(r).Get(m, name)
                                               m *SqliteStore
}

func (s *Registry) Get(store Store, name string) (session *Session, err error) {
  session, err = store.New(s.request, name)
  session := gsessions.NewSession(m, name)

从session.session中获取gsession中Session 结构中的Values map 中的count的值
s.store Store->gsessions.Store

//--------------------------------------------------------------------
session.Set("count", count)

func (s *session) Save() error {
  
  gorilla/session
  func (s *Session) Save(r *http.Request, w http.ResponseWriter) error {
    return s.store.Save(r, w, s)
  }

#154
  func (m *SqliteStore) Save(r *http.Request, w http.ResponseWriter, session *gsessions.Session) error {
    //插入到数据库中，sessionID = 数据库中ID的当前最大值
    if err = m.insert(session); err != nil { 
     
    // 加密session.ID，   
    encoded, err := securecookie.EncodeMulti(session.Name(), session.ID, m.Codecs...)
    // 设置浏览器的cookie
    http.SetCookie(w, gsessions.NewCookie(session.Name(), encoded, session.Options))

  响应response后，最后又来到了中间件  ，？？？？
  r.Use(sessions.Sessions("mysession", store))
       c.Set(DefaultKey, s)

//--------------------------------------------------------------------
再一次时，开始也是新建session，后来发现有cookie，
就解密cookie，得到sessionID，再根据sessionID从数据库中读取并更新session

sessionID！= “” 就使用save()
func (m *SqliteStore) save(session *gsessions.Session) error {
	_, updErr := m.stmtUpdate.Exec(encoded, createdOn, expiresOn, session.ID)


//--------------------------------------------------------------------
  https://blog.csdn.net/wdy_yx/article/details/68059154
  key的更换

  在项目中可能会遇到要更换key的需求，sessions也可以方便的实现
  
  var store = sessions.NewCookieStore(
      []byte("new-authentication-key"),
      []byte("new-encryption-key"),
      []byte("old-authentication-key"),
      []byte("old-encryption-key"),
  )
  
  store通过这样的创建方法，新的sessions会通过一个key pair进行创建，旧的sessions则通过老的key pair读取。
  
  
https://blog.csdn.net/wdy_yx/article/details/67638857
  var hashKey = []byte("1111111111111111")
var blockKey = []byte("1111111111111111")
var sc = securecookie.New(hashKey, blockKey)

// func init() {
//  codecs := securecookie.CodecsFromPairs(
//      []byte("new-hash-key"),
//      []byte("new-block-key"),
//      []byte("old-hash-key"),
//      []byte("old-block-key"),
//  )