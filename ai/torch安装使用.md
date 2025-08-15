# Torch 安装使用

## 安装

### 安装 torch

```
pip install torch torchvision torchaudio
```

### 安装 torchsummary

```
pip install torchsummary
```

### 安装 GPU 版本

1. 安装升级显卡驱动
   1. 命令行输入 nvidia-smi 命令查看显卡驱动版本、支持的CUDA版本、显存大小，如果支持的 CUDA 版本太低则升级显卡驱动
   2. 下载对应版本的显卡驱动，选择 NVIDIA Studio 驱动程序 https://cn.download.nvidia.cn/Windows/580.97/580.97-desktop-win10-win11-64bit-international-nsd-dch-whql.exe
   3. 升级完显卡后再次在终端输入 nvidia-smi 命令查看显卡驱动版本是否升级成功，查看支持的 CUDA 版本

2. 安装 cuda 工具包 
   https://developer.download.nvidia.cn/compute/cuda/13.0.0/local_installers/cuda_13.0.0_windows.exe
   
3. 安装 cuDNN 库 
   https://developer.download.nvidia.cn/compute/cudnn/secure/8.9.7/local_installers/12.x/cudnn-windows-x86_64-8.9.7.29_cuda12-archive.zip
4. 安装 torch 库 

   1. 卸载 cpu 版本
```shell
pip uninstall torch torchvision torchaudio -y
```

   2. 安装 gpu 版本

   我的显卡是 GTX 1050TI的，pytorch 版本最高支持到 2.6，支持的 CUDA 版本是 12.6
   来源：https://github.com/pytorch/pytorch/releases/tag/v2.8.0

```shell
pip install torch==2.6.0 torchvision==0.21.0 torchaudio==2.6.0 --index-url https://download.pytorch.org/whl/cu126
```
   3. 验证安装
```shell
python -c "import torch; print(torch.version.cuda); print(torch.cuda.is_available())"
```

```python
import torch
print(torch.version.cuda);
print(torch.cuda.is_available())

x = torch.rand(5, 3)
print(x)
```

