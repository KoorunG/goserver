package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() { // 단 한번만 호출
		userRouterInstance = &userRouter{
			router: router,
		}
		router.registerPOST("/", userRouterInstance.create)
		router.registerGET("/", userRouterInstance.get)
		router.registerPUT("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)

	})
	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create user")
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get user")
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update user")
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete user")
}
