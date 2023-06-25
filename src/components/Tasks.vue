<script>
import {mapGetters} from "vuex";
import IconArrowLeft from "@/components/icons/IconArrowLeft.vue";
import IconArrowRight from "@/components/icons/IconArrowRight.vue";
import Table from "@/components/Table.vue";

export default {
  name: "Tasks",
  components: {Table, IconArrowRight, IconArrowLeft},
  computed: {
    ...mapGetters({
      synthesis: "synthesizer/getSynthesis",
      timer: "synthesizer/getTimer",
      amountCompleted: "synthesizer/getSynthesisCompleted",
      currentNucleotide: "synthesizer/getNucleotide",
      status: "synthesizer/getStatus"
    }),
  }
}
</script>

<template>
  <div class="synthesis__state">
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
    <div class="tasks">
      <Table title="High" leftPriority="low" rightPriority="medium" :synthesis="synthesis.synthesisHigh"/>
      <Table title="Medium" leftPriority="high" rightPriority="low" :synthesis="synthesis.synthesisMedium"/>
      <Table title="Low" leftPriority="medium" rightPriority="high" :synthesis="synthesis.synthesisLow"/>
    </div>
  </div>
</template>
