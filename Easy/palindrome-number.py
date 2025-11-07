def isPalindrome(x):
        """
        :type x: int
        :rtype: bool
        """

        num = x
        y = 0
        if x < 0:
            x = -x
        
        while x != 0:
            y = (y * 10) + (x % 10)
            x = x // 10


        if num == y:
            return True
        else:
            return False

print(isPalindrome(121))


