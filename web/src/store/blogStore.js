export const blogStore =  {
    state: {
        blogs: []
    },
    getters: {
        blogCount(state) {
            return state.blogs.length
        },
        blogs(state) {
            return state.blogs
        }
    },
    mutations: { // store.commit('setBlogs', blogs) 触发函数的执行
        setBlogs(state, blogs) {
            state.blogs = blogs
        },
        setSingleBlog(state, blog) {
            state.blogs.push(blog)
            console.log('setSingleBlog', blog, state.blogs)
        }
    },
    actions: {
    },
    modules: {
    }
}
