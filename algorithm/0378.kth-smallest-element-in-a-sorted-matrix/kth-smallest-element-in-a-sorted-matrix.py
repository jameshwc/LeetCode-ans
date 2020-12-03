class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        import heapq
        h = []
        for i in range(len(matrix)):
            for j in range(len(matrix[i])):
                heapq.heappush(h, -1 * matrix[i][j])
                if len(h) == k + 1:
                    heapq.heappop(h)
        return -1 * heapq.heappop(h)