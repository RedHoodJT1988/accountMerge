public static void miniMaxSum(List<Integer> arr) {
// Write your code here
    Long min=0L, max=0L, sum=0L;
    for(int i = 0; i< arr.size(); i++){
        for(int j=0; j<arr.size(); j++){
            if(j!=i){
                sum+=arr.get(j);
            }
        }

        if (sum<min || min==0L) min=sum;
        if (sum>max) max=sum;
        sum = 0L;
    }
        System.out.println(min + " " + max);
}