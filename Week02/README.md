我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


思路：
    当在dao层遇到sql.ErrNoRows ，直接使用Wrap往上抛，在servcie层对错误进行判断sql.ErrNoRows，并添加一些辅助信息，最终在调用层打印日志。