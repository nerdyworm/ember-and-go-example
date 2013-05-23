(function() {

  var App = Ember.Application.create();

  DS.Store.create({
    revision: 12,
    adapter: DS.RESTAdapter.create({
      namespace: 'api'
    })
  });

  App.Router.map(function() {
    this.route('create');
  });

  App.Kitten = DS.Model.extend({
    name: DS.attr('string'),
    picture: DS.attr('string')
  });

  App.IndexRoute = Ember.Route.extend({
    model: function() {
      return App.Kitten.find();
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
      }.bind(this));
    }
  });
})();
