$(function() {
  'use strict';
  
  var app = new Vue({
    el: 'body',
    data: {
      todos: []
    },
    methods: {
      load: function () {
        console.log('load', this.text);
        
        $.get('/list', function(data) {
          if (data && data.Status === 'succeeded') {
            this.todos = data.Todos || [];
          }
        }.bind(this));
      },
      
      add: function() {
        console.log('add', this.text);
        
        if (!this.text) return;
        
        var params = { text: this.text };
        this.text = '';
        
        $.post('/add', params, function(data) {
          if (data && data.Status === 'succeeded') {
            this.todos = data.Todos || [];
          }
        }.bind(this));
      },
      
      remove: function (index) {
        console.log('remove', index);
        
        var params = { index: index };
        $.post('/remove', params, function (data) {
          if (data && data.Status === 'succeeded') {
            this.todos = data.Todos || [];
          }
        }.bind(this));
      },
      
      removeAll: function () {
        console.log('removeAll');
        
        $.post('/removeAll', function (data) {
          if (data && data.Status === 'succeeded') {
            this.todos = [];
          }
        }.bind(this));
      }
    },
    created: function() {
     console.log('created'); 
    },
    ready: function() {
      console.log('ready');
      $(this.$el).find('#todo-text').focus();
      this.load();
    },
    filters: {
      date: function (date) {
        return moment(date).format('YYYY/MM/DD HH:mm:ss');
      }
    }
  });
});