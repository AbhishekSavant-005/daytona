from django.utils.deprecation import MiddlewareMixin
from django.core.cache import cache

class RequestCountMiddleware(MiddlewareMixin):
    def process_request(self, request):
        # Increment the total request count in cache
        total_requests = cache.get('total_requests', 0)
        cache.set('total_requests', total_requests + 1)
