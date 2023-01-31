package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/harmannkibue/golang_gin_clean_architecture/internal/usecase"
	db "github.com/harmannkibue/golang_gin_clean_architecture/internal/usecase/repositories"
	"github.com/harmannkibue/golang_gin_clean_architecture/pkg/logger"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"net/http"
)

type ledgerRoute struct {
	u usecase.LedgerUseCase
	l logger.Interface
}

func newVirtualAccountsRoute(handler *gin.RouterGroup, t usecase.LedgerUseCase, l logger.Interface) {
	r := &ledgerRoute{t, l}

	h := handler.Group("/blogs")
	{
		h.POST("/create-blog/", r.createBlog)
		h.GET("/", r.blogs)
		h.GET("/:id", r.blog)

	}
}

type singleBlogResponse struct {
	blog db.Blog `json:"blog"`
}

// @Summary     Fetch single blog by ID
// @Description Show a single blog registered
// @ID          Single blog
// @Tags  	    Blogs
// @Accept      json
// @Produce     json
// @Param        id   path      string  true  "blog ID"
// @Success     200 {object} singleBlogResponse
// @Failure     400 {object} httputil.HTTPError
// @Router      /blogs/{id} [get]
func (route *ledgerRoute) blog(ctx *gin.Context) {
	id := ctx.Param("id")

	blog, err := route.u.GetBlog(ctx, id)
	if err != nil {
		route.l.Error(err, "http - v1 - getting single bank")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

type createBlogRequestBody struct {
	Description string `json:"description"`
}

type createBlogResponse struct {
	Description string `json:"message"`
	CreatedAt   string `json:"created_at"`
}

type listBlogsResponse struct {
	Blogs []createBlogResponse `json:"blogs"`
}

// @Summary     Create a blog
// @Description Create a blog
// @ID          Create a blog
// @Tags  	    Blogs
// @Accept      json
// @Produce     json
// @Param       request body createBlogRequestBody true "Create blog request body"
// @Success     201 {object} createBlogResponse
// @Failure     400 {object} httputil.HTTPError
// @Failure     500 {object} httputil.HTTPError
// @Router      /blogs/create-blog/ [post]
func (route *ledgerRoute) createBlog(ctx *gin.Context) {
	var body createBlogRequestBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		route.l.Error(err, "http - v1 - create a blog route")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	blog, err := route.u.CreateBlog(ctx, body.Description)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	ctx.JSON(http.StatusCreated, blog)
}

// @Summary     List all the Blogs
// @Description Show all blogs registered
// @ID          Fetch Blog
// @Tags  	    Blogs
// @Accept      json
// @Produce     json
// @Param 		Page query string false "1" "The page number that you want items for"
// @Param 		ItemsPerPage query string false "10" "The number of items per page"
// @Success     200 {object} listBlogsResponse
// @Failure     400 {object} httputil.HTTPError
// @Router      /blogs/ [get]
func (route *ledgerRoute) blogs(ctx *gin.Context) {

	page := ctx.Request.URL.Query().Get("Page")
	limit := ctx.Request.URL.Query().Get("ItemsPerPage")

	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}

	blogs, err := route.u.ListBlogs(ctx, usecase.ListBlogsParams{
		Page:  page,
		Limit: limit,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, blogs)
}
