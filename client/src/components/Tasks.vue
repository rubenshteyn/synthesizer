<script>
import {mapGetters} from "vuex";
import IconArrowLeft from "@/components/icons/IconArrowLeft.vue";
import IconArrowRight from "@/components/icons/IconArrowRight.vue";
import Table from "@/components/Table.vue";

export default {
  name: "Tasks",
  data() {
    return {
      priorities: [
        {
          name: "synthesisHigh",
          title: "High",
          leftPriority: "synthesisLow",
          rightPriority: "synthesisMedium",
        },
        {
          name: "synthesisMedium",
          title: "Medium",
          leftPriority: "synthesisHigh",
          rightPriority: "synthesisLow"
        },
        {
          name: "synthesisLow",
          title: "Low",
          leftPriority: "synthesisMedium",
          rightPriority: "synthesisHigh"
        }
      ]
    }
  },
  components: {Table, IconArrowRight, IconArrowLeft},
  computed: {
    ...mapGetters({
      synthesis: "synthesizer/getSynthesis"
    })
  }
}
</script>

<template>
  <div class="synthesis__state">
    <div class="tasks">
      <Table v-for="priority in priorities" :key="priority.name"
             :title="priority.title"
             :leftPriority="priority.leftPriority"
             :rightPriority="priority.rightPriority"
             :synthesis="synthesis[priority.name]"
             :queue="priority.name"/>
    </div>
  </div>
</template>
