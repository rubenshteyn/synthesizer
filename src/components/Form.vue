<script>
import {v4 as uuidv4} from "uuid";
import {mapActions} from "vuex";

export default {
  name: "Form",
  data() {
    return {
      formula: null,
      priority: "medium",
      error: null
    }
  },
  methods: {
    ...mapActions({
      "create": "synthesizer/addSynthesisMethod",
      "updateTimer": "synthesizer/updateTimerMethod",
    }),
    createTask() {
      let form = {
        id: uuidv4(),
        nucleotides: this.formula.toLowerCase().split(""),
        priority: this.priority,
        work: false
      }
      this.create(form)
      this.updateTimer()
    },

    addPriority(priority) {
      this.priority = priority
    },
  },
  computed: {
    isChanged() {
      const regular = /^([a, t, g, c]+|\d+)$/i;
      if (!this.formula) {
        this.error = null
        return false
      }

      if (!regular.test(this.formula)) {
        this.error = "Введите корректную формулу!"
        return true
      }

      if (this.formula.length < 6) {
        this.error = "Слишком короткая формула!"
        return true
      }

      if (this.formula.length > 120) {
        this.error = "Слишком длинная формула!"
        return true
      }

      if (!this.priority) {
        this.error = "Выберите приоритет!"
        return true
      }

      if (!this.formula || !regular.test(this.formula) || !this.priority) {
        return true
      }
    },
  }
}

</script>

<template>
  <div class="synthesis__form">
    <div class="synthesis__block">
      <p v-if="isChanged" class="error--formula">{{ this.error }}</p>
      <input
          placeholder="Введите формулу"
          v-model="formula"
          type="text"
          class="synthesis__description"
          :class="isChanged ? 'error--border' : ''"
      />
      <button
          :disabled=isChanged
          @click="createTask()"
          class="synthesis__btn submit">
        Синтез
      </button>
    </div>
    <div class="synthesis--priority">
      <p class="priority--title">Выберите приоритет:</p>
      <div class="synthesis__btns">
        <button
            @click="addPriority('high')"
            :class="priority === 'high' ? 'synthesis__btn high opacity' : 'synthesis__btn high'">
          Высокий
        </button>
        <button
            @click="addPriority('medium')"
            :class="priority === 'medium' ? 'synthesis__btn medium opacity' : 'synthesis__btn medium'">
          Средний
        </button>
        <button
            @click="addPriority('low')"
            :class="priority === 'low' ? 'synthesis__btn low opacity' : 'synthesis__btn low'">
          Низкий
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>