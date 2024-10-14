<script>
import {mapActions, mapGetters} from "vuex";
import ModalService from "@/components/modals/ModalService.vue";
import IconArrowRight from "@/components/icons/IconArrowRight.vue";
import IconArrowLeft from "@/components/icons/IconArrowLeft.vue";
import Nucleotides from "@/components/Nucleotides.vue";
import Form from "@/components/Form.vue";
import Tasks from "@/components/Tasks.vue";
import StateInfo from "@/components/StateInfo.vue";

export default {
  data() {
    return {
      counter: 0,
      amountWorks: 0
    }
  },
  components: {
    StateInfo,
    Tasks,
    Form,
    Nucleotides,
    IconArrowLeft,
    IconArrowRight,
    ModalService,
  },

  mounted() {
    this.updateTimer()
    this.work()
  },
  methods: {
    ...mapActions({
      "delete": "synthesizer/deleteSynthesisMethod",
      "statusUpdate": "synthesizer/updateStatusSynthesisMethod",
      "updateTimer": "synthesizer/updateTimerMethod",
      "subtractionTime": "synthesizer/subtractionTimeMethod",
      "updateSynthesisCompleted": "synthesizer/updateSynthesisCompletedMethod",
      "updateCurrentNucleotide": "synthesizer/updateCurrentNucleotideMethod",
      "updateStatus": "synthesizer/updateStatusMethod"
    }),

    timerIteration(priority) {
      const isSynthes = this.synthesis[priority][0]
      if (isSynthes
          && isSynthes.length !== 0
          && !isSynthes.work
          && this.status !== "На обслуживании"
          && this.status !== "Редактирование"
          && this.amountWorks === 0) {

        const data = {
          synthes: {
            id: this.synthesis[priority][0].id,
            work: true,
            nucleotides: this.synthesis[priority][0].nucleotides
          },
          queue: priority
        }
        this.updateStatus("Занят")
        this.statusUpdate(data)
        this.amountWorks++

        isSynthes.nucleotides.forEach((el, i) => {
          setTimeout(() => {
            this.updateCurrentNucleotide(el)
            this.subtractionTime()
          }, 1000 * (i + 1))
        })

        setTimeout(() => {
          this.updateCurrentNucleotide(null)
          this.delete({synthes: this.synthesis[priority][0], queue: priority})
          this.updateSynthesisCompleted()
          this.updateTimer()
          this.counter = 0
          this.amountWorks = 0
          this.work()
        }, isSynthes.nucleotides.length * 1000)
      }
    },

    work() {
      const conditionCompleted = this.amountCompleted !== 0 && this.amountCompleted % 5 === 0
      if (conditionCompleted && this.status === "Занят" && this.counter < 2) {
        this.counter++
        this.updateStatus("На обслуживании")
      }
      if (this.status === "На обслуживании" && this.counter === 1) {
        this.counter++
        setTimeout(() => {
          this.updateStatus("Бездействует")
        }, 5000)
      }
      if (this.status !== "На обслуживании") {
        const totalLength = this.synthesis.synthesisHigh.length === 0 && this.synthesis.synthesisMedium.length === 0 && this.synthesis.synthesisLow.length === 0
        if (this.synthesis.synthesisHigh.length !== 0) {
          this.timerIteration("synthesisHigh")
        }
        if (this.synthesis.synthesisHigh.length === 0) {
          this.timerIteration("synthesisMedium")
        }
        if (this.synthesis.synthesisHigh.length === 0 && this.synthesis.synthesisMedium.length === 0) {
          this.timerIteration("synthesisLow")
        }
        if (totalLength && this.amountCompleted !== 4) {
          this.updateStatus("Бездействует")
        }
        return true
      }
    }
  },
  computed: {
    ...mapGetters({
      synthesis: "synthesizer/getSynthesis",
      amountCompleted: "synthesizer/getSynthesisCompleted",
      status: "synthesizer/getStatus"
    }),
  }
}
</script>

<template>
  <main>
    <div v-if="work()" class="synthesizer">
      <div class="synthesizer__block state--block">
        <Nucleotides/>
        <StateInfo />
      </div>
      <div class="synthesizer__block">
        <Form/>
        <Tasks/>
      </div>
    </div>
    <ModalService v-else/>
  </main>
</template>

