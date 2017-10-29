<template>
  <div>
    <vue-good-table
      title="Devices Available"
      :columns="columns"
      :rows="rows"
      :paginate="true"
      :lineNumbers="true"/>
  </div>
</template>

<script>
  export default {
    name: 'test',
    data () {
      return {
        columns: [
          {
            label: 'Name',
            field: 'name',
            filterable: true
          },
          {
            label: 'Address IPv4',
            field: 'addressIP4',
            filterable: true
          },
          {
            label: 'Address IPv6',
            field: 'addressIP6',
            filterable: true
          }
        ],
        rows: []
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

                self.rows.push({'name': device.Name, 'addressIP4': addressIP4, 'addressIP6': addressIP6})
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
