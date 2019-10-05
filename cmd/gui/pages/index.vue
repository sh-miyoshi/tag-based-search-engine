<template>
  <div class="container">
    <div>
      <head>
        Logoとか
      </head>

      <div>
        <h3>タグリスト</h3>
        タグ一覧(追加・削除)
      </div>

      <div>
        <h3>コンテンツ一覧</h3>
        写真とか動画とか・・・
        <br />
        {{ this.$store.getters.getData }}
      </div>
    </div>
  </div>
</template>

<script>
export default {}
</script>

<style></style>
