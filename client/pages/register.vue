<template>
  <v-card>
    <v-card-title>近畿大学電子計算機研究会 入部フォーム</v-card-title>
    <v-card-text>
      <v-alert type="error" v-if="error">
        {{ error }}
      </v-alert>
      <v-form ref="register_form">
        <v-text-field
          v-model="id"
          label="ユーザーID (4文字~15文字, 英数字)"
          :rules="[required, id_length, alpha_num_check]"
        />
        <v-text-field
          v-model="password"
          label="パスワード (8文字以上)"
          type="password"
          :rules="[required, password_length]"
        />
        <v-text-field
          v-model="passwordconfirm"
          label="パスワード (確認)"
          type="password"
          :rules="[required, password_confirm]"
        />
        <v-text-field
          v-model="studentnumber"
          label="学籍番号 (例 2010370300)"
          :rules="[required, number_check]"
        />
        <v-text-field
          v-model="name"
          label="氏名"
          :rules="[required]"
        />
        <v-text-field
          v-model="mailaddress"
          label="メールアドレス"
          :rules="[required, mail_check]"
        />
      </v-form>
    </v-card-text>
    <v-card-actions>
      <v-btn text v-on:click="submit">送信する</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  data() {
      return {
        id: "",
        password: "",
        passwordconfirm: "",
        studentnumber: "",
        name: "",
        mailaddress: "",
        error: "",
        required: value => !!value || "必須項目です",
        id_length: value => (value.length >= 4 && value.length <= 15) || "4文字以上15文字以内で入力してください",
        password_length: value => value.length >= 8 || "8文字以上で入力してください",
        password_confirm: value => {
          if (this.password !== value) {
            return "パスワードが一致しません"
          }
        },
        number_check: value => !isNaN(value) || "半角数字のみで入力してください",
        alpha_num_check: value => {
          if (!value.match(/^[A-Za-z0-9]*$/))
            return "半角英数字のみで入力してください"
        },
        mail_check: value => {
          const studentNumberReg = new RegExp(this.studentnumber+".@kindai.ac.jp")
          if (!value.match(studentNumberReg) || !value.match(/^[A-Za-z0-9@.+]*$/))
            return "大学から交付されたメールアドレスを入力してください"
        }
      }
  },
  head() {
    return {
      title: '新規登録',
    }
  },
  methods: {
    submit() {
      if (this.$refs.register_form.validate()) {
        const params = new URLSearchParams()
        params.append('Token', this.$route.query.token)
        params.append('ID', this.id)
        params.append('Password', this.password)
        params.append('Name', this.name)
        params.append('EmailAddress', this.mailaddress)
        params.append('StudentNumber', this.studentnumber)
        this.$axios.$post('/api/register', params)
          .then((result) => {
            console.log(result)
            this.$router.push('/thanks')
          })
          .catch((e) => {
            if (e.response) {
              this.error = e.response.data.Msg
            } else {
              this.error = "予期せぬなエラーが発生しました. 問い合わせてください."
            }
          })
      } else {
        console.log("out")
      }
    }
  }
}
</script>