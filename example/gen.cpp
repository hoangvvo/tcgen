#include <iostream>
#include <fstream>
#include <string>
#include <random>
using namespace std;

std::random_device myRandomDevice;
unsigned seed = myRandomDevice();
default_random_engine myRandomEngine(seed);

int randomBetweenTwoNum(int min, int max)
{
  uniform_int_distribution<int> myUnifIntDist(min, max);
  return myUnifIntDist(myRandomEngine);
}

int main(int argc, char **argv)
{
  ofstream ofs;
  ofs.open("SUM.INP", ifstream::out);

  ofs << randomBetweenTwoNum(0, 10000) << " " << randomBetweenTwoNum(0, 10000);

  ofs.close();

  return 0;
}