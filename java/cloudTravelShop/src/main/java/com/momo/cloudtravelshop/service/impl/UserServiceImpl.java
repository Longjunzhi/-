package com.momo.cloudtravelshop.service.impl;

import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.momo.cloudtravelshop.entity.User;
import com.momo.cloudtravelshop.mapper.UserMapper;
import com.momo.cloudtravelshop.service.UserService;
import org.springframework.stereotype.Service;

@Service
public class UserServiceImpl extends ServiceImpl<UserMapper, User> implements UserService {
}
