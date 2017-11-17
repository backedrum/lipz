<template>
    <vue-good-table
      title="Captured packets"
      :columns="columns"
      :rows="rows"
      :paginate="true"
      :globalSearch="true"
      styleClass="table table-bordered table-striped"
      :lineNumbers="true">
      <div slot="emptystate">
        Please wait, capturing packets...
      </div>
    </vue-good-table>
</template>

<script>
  export default {
    name: 'packetsTable',
    data () {
      return {
        columns: [
          {
            label: 'Protocol',
            field: 'protocol'
          },
          {
            label: 'Src IP',
            field: 'srcIP'
          },
          {
            label: 'Src port',
            field: 'srcPort'
          },
          {
            label: 'Dst IP',
            field: 'dstIP'
          },
          {
            label: 'Dst port',
            field: 'dstPort'
          },
          {
            label: 'Payload',
            field: 'payload'
          }
        ],
        rows: []
      }
    },
    methods: {
      getTableData () {
        var self = this

        setTimeout(function () {
          self.$http.get('/api/capture/' + self.$route.params.interfaceName)
            .then(response => {
              for (var i = 0; i < response.body.NetPacketInfoList.length; i++) {
                var packet = response.body.NetPacketInfoList[i]
                self.rows.push({'protocol': packet.Protocol, 'srcIP': packet.Src, 'srcPort': packet.SrcPort, 'dstIP': packet.Dst, 'dstPort': packet.DstPort, 'payload': packet.Payload})
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
