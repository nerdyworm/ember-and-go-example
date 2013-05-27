(function() {

  var App = window.App = Ember.Application.create();

  DS.Store.create({
    revision: 12,
    adapter: DS.RESTAdapter.create({
      namespace: 'api'
    })
  });

  App.Router.map(function() {
    this.route('create');
    this.route('edit', {path: '/edit/:kitten_id'});
  });

  App.Kitten = DS.Model.extend({
    name: DS.attr('string'),
    picture: DS.attr('string')
  });

  App.IndexRoute = Ember.Route.extend({
    model: function() {
      return App.Kitten.find();
    },

    events: {
      deleteKitten: function(kitten) {
        kitten.deleteRecord();
        kitten.save();
      }
    }
  });

  App.CreateController = Ember.Controller.extend({
    name: null,
    save: function() {
      var kitten = App.Kitten.createRecord({
        name: this.get('name')
      });

      kitten.save().then(function() {
        this.transitionToRoute('index');
        this.set('name', '');
      }.bind(this));
    }
  });

  App.EditController = Ember.ObjectController.extend({
    save: function() {
      var kitten = this.get('model');
      kitten.save().then(function() {
        this.transitionToRoute('index');
      }.bind(this));
    }
  });
})();
