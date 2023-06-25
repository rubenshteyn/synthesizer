import { createStore } from 'vuex'
import synthesizer from './modules/synthesizer'
const store = createStore({
    state: {  },
    mutations: {  },
    actions: {  },
    modules: { synthesizer: synthesizer }
})


export default store