from rest_framework import serializers
from .models import Collection, Movie

from rest_framework import serializers

class MovieSerializer(serializers.ModelSerializer):
    class Meta:
        model = Movie
        fields = ['external_id', 'title', 'description', 'genre']
        extra_kwargs = {
            'external_id': {'required': True},
            'title': {'required': True},
            'description': {'required': True},
            'genre': {'required': True},
        }

class CollectionSerializer(serializers.ModelSerializer):
    class Meta:
        model = Collection
        fields = ['name']
        extra_kwargs = {
            'name': {'required': True},
        }
