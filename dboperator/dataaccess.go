package dboperator

// ConnectionInfoContainer 定义了连接字符串的访问功能
type ConnectionInfoContainer interface {
	SetConnectionString(connectionStr string)
	GetConnectionString() string
}

// DBConnection 数据库连接接口，定义了打开连接，关闭连接以及获取当前连接状态的功能
type DBConnection interface {
	Open()
	Close()
	GetState() int8
}

// DBCommand 数据库命令接口，实现基本的设置连接和执行功能
type DBCommand interface {
	SetConnection(conn *DBConnection)
	ExecuteNonQuery(sql string)
}

// ConnectionInfo 包含基本的数据库连接字符串
type ConnectionInfo struct {
	connectionString string
}

// SetConnectionString 实现将参数传入的connectionStr保存到接收对象指针的对应成员connectionString中
func (ci *ConnectionInfo) SetConnectionString(connectionStr string) {
	ci.connectionString = connectionStr
}

// GetConnectionString 返回接收对象指针的connectionString的值
func (ci *ConnectionInfo) GetConnectionString() string {
	return ci.connectionString
}
