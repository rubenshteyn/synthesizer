<script>
import {mapGetters} from "vuex";

export default {
  name: "StateInfo",
  computed: {
    ...mapGetters({
      synthesis: "synthesizer/getSynthesis",
      timer: "synthesizer/getTimer",
      amountCompleted: "synthesizer/getSynthesisCompleted",
      status: "synthesizer/getStatus"
    })
  }
}
</script>

<template>
  <div class="synthes--block-state">
    <div class="synthes--el">
      <p class="synthesis__state--description">Состояние синтезатора: <span
          class="description--data">{{ status }}</span></p>
      <p class="synthesis__state--description">Синтезатор закончит работу через:
        <span class="description--data">{{ timer }}с</span>
      </p>
    </div>
    <div class="synthes--el">
      <p class="synthesis__state--description">Количество выполненных задач:
        <span class="description--data">
          {{ amountCompleted }}</span>
      </p>
      <p class="synthesis__state--description">Количество задач в очереди:
        <span class="description--data">
          {{
            synthesis.synthesisHigh.length + synthesis.synthesisMedium.length + synthesis.synthesisLow.length
          }}</span>
      </p>
    </div>
  </div>
</template>
