<template>
  <div>
    <v-table
      :width="800"
      :height="350"
      :columns="tableConfig.columns"
      :table-data="tableConfig.tableData"
      even-bg-color="#f4f4f4"
      row-hover-color="#eee"
      row-click-color="#edf7ff"
    ></v-table>
  </div>
</template>

<script>
  export default{
    name: 'devices',
    data () {
      return {
        tableConfig: {
          tableData: [],
          columns: [
            {field: 'name', title: 'Name', width: 150, titleAlign: 'center', columnAlign: 'center', isFrozen: true},
            {field: 'addressIP4', title: 'Address IP4', width: 430, titleAlign: 'center', columnAlign: 'center', isFrozen: true},
            {field: 'addressIP6', title: 'Address IP6', width: 280, titleAlign: 'center', columnAlign: 'center', isFrozen: true}
          ]
        }
      }
    },
    methods: {
      getTableData () {
        var self = this

        setTimeout(function () {
          self.$http.get('/api/devices/list')
            .then(response => {
              for (var i = 0; i < response.body.Devices.length; i++) {
                var device = response.body.Devices[i]
                var addressIP6 = device.Addresses.length > 0
                  ? device.Addresses[0].IP : '<undef>'
                var addressIP4 = device.Addresses.length > 1
                  ? device.Addresses[1].IP : '<undef>'

                self.tableConfig.tableData.push({'name': device.Name, 'addressIP4': addressIP4, 'addressIP6': addressIP6})
              }
            }, response => {
              console.log(response)
            })
        }, 100)
      }
    },
    created () {
      this.getTableData()
    }
  }
</script>
