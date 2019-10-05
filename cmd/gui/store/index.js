import jsonData from '~/database.json'

export const state = () => ({
  data: jsonData
})

export const getters = {
  getData: (state) => {
    return state.data
  }
}
