<template>
  <EmptyLayout>
    <div class="login-box">
      <div class="login-logo">
        <a href="./">LOGIN</a>
      </div>
      <div class="login-box-body">
        <p class="login-box-msg"></p>

        <form @submit.prevent="doLogin">
          <input type="hidden" name="callback" value="{[{.callback}]}">
          <div class="form-group has-feedback">
            <input type="text" class="form-control" placeholder="帐号" v-model="form.User">
            <span class="glyphicon glyphicon-user form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="password" class="form-control" placeholder="密码" v-model="form.Password">
            <span class="glyphicon glyphicon-lock form-control-feedback"></span>
          </div>
          <div class="row">
            <div class="col-xs-8">
              <div class="checkbox icheck">
                <label>
                  <input type="checkbox" v-model="form.Remember"> 记住密码</label>
              </div>
            </div>
            <div class="col-xs-4">
              <button type="submit" class="btn btn-primary btn-block btn-flat">登陆</button>
            </div>
          </div>
        </form>

        <div class="social-auth-links text-center">
          <p>- OR -</p>
          <aa href="/oauth2/login?callback={[{.callback}]}" class="btn btn-block btn-success btn-flat"><i
              class="fa fa-lock"></i> Sign in using
            OAuth2</aa>
        </div>

        <a href="/">返回首页</a><br>
      </div>
    </div>
  </EmptyLayout>
</template>

<script>
  import EmptyLayout from "../layout/empty.vue";

  export default {
    components: {
      EmptyLayout
    },
    data() {
      return {
        form: {
          User: "admin",
          Password: "admin",
          Remember: true,
        }
      }
    },
    props: {},
    created() { },
    mounted() {
      $('input')
        .iCheck({
          checkboxClass: 'icheckbox_square-blue',
          radioClass: 'iradio_square-blue',
          increaseArea: '20%' // optional
        })
        .on('ifChanged', event => {
          this.form.Remember = event.target.checked
        });
    },
    destroyed() { },
    methods: {
      doLogin(o) {
        this.$post("/admin/login", this.form, x => {
          if (x.no != 0) return alert(x.data);
          $(".login-box-msg").html("登陆成功");
          this.$router.push('/admin/dashboard');
        });
      }
    },
    filters: {},
    watch: {},
    computed: {},
  }
</script>