import unittest
from main import solution

class TestMain(unittest.TestCase):
    def test_solution(self):
        self.assertEqual(solution("test 5 a0A pass007 ?xy1"), 7)
        self.assertEqual(solution("test11569 5 a0A pass007 @????@@@@@@@@"), 9)
        self.assertEqual(solution("5"), 1)
        self.assertEqual(solution("a"), -1)
        self.assertEqual(solution("@@@???"), -1)

if __name__ == '__main__':
    unittest.main()