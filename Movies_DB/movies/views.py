from django.shortcuts import get_object_or_404, render
from rest_framework.decorators import api_view, permission_classes
from rest_framework.response import Response
from rest_framework import status, viewsets
from django.core.cache import cache
from rest_framework.permissions import AllowAny
from django.contrib.auth.models import User
import requests
from .models import Movie, Collection
from .serializers import CollectionSerializer, MovieSerializer
from .services import fetch_movies_from_api

### MOVIE HANDLING ###

# Function to fetch movies from the external API



# View to list movies from both the API and the database

@api_view(['GET'])
@permission_classes([AllowAny])
def list_movies(request):
    movies_from_api = fetch_movies_from_api()

    print(f"Movies fetched from API: {movies_from_api}")  # Debugging output

    if not movies_from_api:
        return Response({"error": "No movies fetched from external API"}, status=status.HTTP_500_INTERNAL_SERVER_ERROR)

    db_movies = Movie.objects.all()

    combined_movies = {
        "api_movies": movies_from_api,
        "db_movies": [
            {"external_id": movie.external_id, "title": movie.title, "description": movie.description, "genre": movie.genre}
            for movie in db_movies
        ]
    }

    print(f"Combined Movies: {combined_movies}")  # Debugging output

    return Response(combined_movies, status=status.HTTP_200_OK)



### COLLECTION HANDLING ###




@api_view(['POST'])
@permission_classes([AllowAny])  # Allow access without authentication
def create_collection(request):
    collection_name = request.data.get('name')
    movie_data = request.data.get('movies', [])

    if collection_name is None:
        return Response({'error': 'Collection name is required.'}, status=status.HTTP_400_BAD_REQUEST)

    # Create the new collection
    collection = Collection.objects.create(name=collection_name)

    # Add movies to the collection
    for movie_item in movie_data:
        external_id = movie_item.get('external_id')
        try:
            movie = Movie.objects.get(external_id=external_id)
            collection.movies.add(movie)
        except Movie.DoesNotExist:
            return Response({'error': f'Movie with external_id {external_id} does not exist.'}, status=status.HTTP_404_NOT_FOUND)

    return Response({"message": "Collection created successfully", "collection_id": collection.id}, status=status.HTTP_201_CREATED)


@api_view(['GET'])
def get_collection(request, collection_id):
    try:
        collection = Collection.objects.get(id=collection_id)
        movies = [{"external_id": movie.external_id} for movie in collection.movies.all()]
        return Response({"id": collection.id, "name": collection.name, "movies": movies}, status=status.HTTP_200_OK)
    except Collection.DoesNotExist:
        return Response({'error': 'Collection not found.'}, status=status.HTTP_404_NOT_FOUND)

@api_view(['PUT'])
def update_collection(request, collection_id):
    try:
        collection = Collection.objects.get(id=collection_id)
        collection_name = request.data.get('name')
        movie_data = request.data.get('movies', [])

        if collection_name is not None:
            collection.name = collection_name
        
        # Clear existing movies
        collection.movies.clear()

        # Add new movies
        for movie_item in movie_data:
            external_id = movie_item.get('external_id')
            try:
                movie = Movie.objects.get(external_id=external_id)
                collection.movies.add(movie)
            except Movie.DoesNotExist:
                return Response({'error': f'Movie with external_id {external_id} does not exist.'}, status=status.HTTP_404_NOT_FOUND)

        collection.save()
        return Response({"message": "Collection updated successfully"}, status=status.HTTP_200_OK)
    except Collection.DoesNotExist:
        return Response({'error': 'Collection not found.'}, status=status.HTTP_404_NOT_FOUND)
    
    
@api_view(['DELETE'])
def delete_collection(request, collection_id):
    try:
        collection = Collection.objects.get(id=collection_id)
        collection.delete()
        return Response({"message": "Collection deleted successfully"}, status=status.HTTP_204_NO_CONTENT)
    except Collection.DoesNotExist:
        return Response({'error': 'Collection not found.'}, status=status.HTTP_404_NOT_FOUND)



@api_view(['GET'])
@permission_classes([AllowAny])  # Allow access without authentication
def get_request_count(request):
    total_requests = cache.get('total_requests', 0)
    return Response({"total_requests": total_requests})


### USER AUTHENTICATION ###

@api_view(['POST'])
@permission_classes([AllowAny])
def register(request):
    username = request.data.get('username')
    password = request.data.get('password')

    if username is None or password is None:
        return Response({'error': 'Please provide both username and password.'}, status=status.HTTP_400_BAD_REQUEST)

    # Check if user already exists
    if User.objects.filter(username=username).exists():
        return Response({'error': 'Username already taken.'}, status=status.HTTP_400_BAD_REQUEST)

    # Create new user
    user = User.objects.create_user(username=username, password=password)
    user.save()

    return Response({'message': 'User registered successfully'}, status=status.HTTP_201_CREATED)


### MOVIE VIEWSET ###

class MovieViewSet(viewsets.ModelViewSet):
    permission_classes = [AllowAny]  # Allow access without authentication
    serializer_class = MovieSerializer

    def get_queryset(self):
        return Movie.objects.all()  # Returns all movies, feel free to customize
    

class CollectionViewSet(viewsets.ModelViewSet):
    permission_classes = [AllowAny]  # Allow access without authentication
    serializer_class = CollectionSerializer

    def get_queryset(self):
        return Collection.objects.all()  # Return all collections