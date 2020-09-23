#include<iostream>
#include<vector>

using namespace std;

queue := make([]string, 0)
	for i := 1; i <= 9; i++ {
		queue = append(queue, strconv.Itoa(i))
	}
	// 位数减1 代表第一层是 [1,2,3,4,5,6,7,8,9]
	n--
	for len(queue) != 0 && n > 0{
		// 获取 当前队列的长度
		l := len(queue)
		// 遍历当前层的数据
		for j := 0; j < l; j++ {
			v := queue[0]
			queue = queue[1:]
			fmt.Print(v, " ")
			//  在其字符串后追加0-9的字符
			for i := 0; i < 10; i++ {
				// 10 11 12 13 14 15 16 17 18 19 20 21.....
				// 100 ....
				// 1000
				queue = append(queue, v+strconv.Itoa(i))
			}
		}
		//位数减一
		n--
	}
void func(int n){
  vector<int> datas;
  for(int i=1;i<=9;i++){
    datas.push_back(i);
  }
  n--;
  while(n>0&&datas.size()>0){
    int size=datas.size();
    for(int i=0;i<size;i++){
      
    }
  }

}