export default {
  loadActiveProducts: `
    query {
          loadActiveProducts{
            SKU
            name
            description
            price
            stock
            startDate
            endDate
            createdAt
          }
    } 
    `
}